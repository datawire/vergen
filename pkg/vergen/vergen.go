package vergen

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Generator struct {
	GitRepo          *git.Repository
	ReleaseTagPrefix string
}

func NewGenerator() (*Generator, error) {
	opts := git.PlainOpenOptions{DetectDotGit: true}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	gitRepo, err := git.PlainOpenWithOptions(cwd, &opts)
	if err != nil {
		return nil, err
	}

	return &Generator{GitRepo: gitRepo}, nil
}

type Version interface {
	String() string
}

type CalendarVersion struct {
	// Year of the version.
	Year int

	// Month of the version.
	Month time.Month

	// Day of the version.
	Day int

	// RevisionID is an offset from 0 indicating the number of releases on the date.
	RevisionID int
}

func (v *CalendarVersion) String() string {
	date := fmt.Sprintf("%d.%02d.%02d", v.Year, v.Month, v.Day)
	return fmt.Sprintf("%s-%d", date, v.RevisionID)
}

type PreviewVersion struct {
	CommitID   string
	Branch     string
	Authority  string
	RevisionID string
}

func (v *PreviewVersion) String() string {
	result := fmt.Sprintf("%s.%s.%s", v.CommitID, v.Branch, v.Authority)

	if v.RevisionID != "" {
		result += fmt.Sprintf("-%s", v.RevisionID)
	}

	return result
}

func (g *Generator) getVersionTagNames(prefix string) ([]string, error) {
	var result []string
	tagRefs, err := g.GitRepo.Tags()
	if err != nil {
		return result, err
	}

	err = tagRefs.ForEach(func(r *plumbing.Reference) error {
		name := r.Name().Short()
		if strings.HasPrefix(name, g.ReleaseTagPrefix) {
			result = append(result, name)
		}

		return nil
	})

	sort.Strings(result)

	return result, nil
}

func getCurrentReleaseVersion(tags []string) string {
	if len(tags) > 0 {
		return tags[len(tags)-1]
	} else {
		return ""
	}
}

func (g *Generator) computeNextRevision(tags []string) int {
	latestRevision := 0

	for i := range tags {
		version := strings.Replace(tags[i], g.ReleaseTagPrefix, "", 1)
		parts := strings.Split(version, "-")
		if len(parts) == 2 {
			rev, err := strconv.Atoi(parts[1])

			// If a person uses only this tool it should not possible to generate a "bad" revision. However, if a person
			// uses this tool and sometimes manually generates version they might create a non-integer revision by
			// mistake. This tool considers such revisions as equivalent to 0.
			if err != nil {
				rev = 0
			}

			if rev > latestRevision {
				latestRevision = rev
			}
		}
	}

	return latestRevision
}

func (g *Generator) NextReleaseVersion() (Version, error) {
	timeUTC := time.Now().UTC()
	date := fmt.Sprintf("%d.%02d.%02d", timeUTC.Year(), timeUTC.Month(), timeUTC.Day())
	prefix := fmt.Sprintf("%s%s", g.ReleaseTagPrefix, date)

	tags, err := g.getVersionTagNames(prefix)
	if err != nil {
		return nil, err
	}

	nextRevision := g.computeNextRevision(tags)

	result := &CalendarVersion{
		Year:       timeUTC.Year(),
		Month:      timeUTC.Month(),
		Day:        timeUTC.Day(),
		RevisionID: nextRevision,
	}

	return result, nil
}

func (g *Generator) LatestReleaseVersion(fallback string) (string, error) {
	tags, err := g.getVersionTagNames(g.ReleaseTagPrefix)
	if err != nil {
		return "", err
	}

	result := getCurrentReleaseVersion(tags)

	if result == "" && fallback != "" {
		result = fallback
	}

	return result, nil
}

// PreviewVersion generates a "preview" style version identifier.
//
// branch - Manually override branch name detection. This is useful in CI where vergen can be used with a detached head
//          and the branch can be passed along via environment variable.
//
// authority - authorized person or system generating the preview version.
//
// generateRevisionID - generate a revision ID to distinguish between two invocations. Useful when used on dirty Git
//                      trees.
func (g *Generator) PreviewVersion(branch, authority string, generateRevisionID bool) (Version, error) {
	result := &PreviewVersion{
		CommitID:  "none",
		Branch:    "none",
		Authority: "unknown",
	}

	if branch != "" {
		result.Branch = branch
	}

	if authority != "" {
		result.Authority = authority
	}

	if generateRevisionID {
		result.RevisionID = generateRevision()
	}

	head, err := g.GitRepo.Head()
	if err == nil {
		result.CommitID = head.Hash().String()
		result.Branch = head.Name().Short()
	}

	result.Branch = normalizeBranchName(result.Branch)

	return result, nil
}

func generateRevision() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
