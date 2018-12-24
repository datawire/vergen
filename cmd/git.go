package cmd

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
	"os"
	"sort"
	"strings"
)

func GetTagNames(r *git.Repository, prefix string) []string {
	var result []string

	tagRefs, err := r.Tags()
	if err != nil {
		log.Fatal(err)
	}

	err = tagRefs.ForEach(func(r *plumbing.Reference) error {
		name := r.Name().Short()
		if strings.HasPrefix(name, prefix) {
			result = append(result, name)
		}

		return nil
	})

	sort.Strings(result)

	return result
}

func GetRepo() *git.Repository {
	opts := git.PlainOpenOptions{DetectDotGit: true}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpenWithOptions(cwd, &opts)
	CheckIfError(err)

	return repo
}
