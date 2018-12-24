package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

func init() {
	cmd := ReleaseVersionCmd()
	RootCmd.AddCommand(cmd)

	cmd.Flags().StringVar(&prefix, "prefix", "release/", "A prefix the comes before the version")
}

func ReleaseVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "next",
		Short: "Examines Git tags to determine the next release version",
		Run:   nextReleaseVersion,
	}
}

func nextReleaseVersion(cmd *cobra.Command, args []string) {
	repo := GetRepo()

	timeUTC := time.Now().UTC()
	date := fmt.Sprintf("%d.%02d.%02d", timeUTC.Year(), timeUTC.Month(), timeUTC.Day())
	prefix := fmt.Sprintf("%s%s", prefix, date)

	tags := GetTagNames(repo, prefix)
	nextRevision := computeNextRevision(tags)

	fmt.Printf("%s-%d\n", date, nextRevision)
}

func computeNextRevision(tags []string) int {
	latestRevision := 0

	for i := range tags {
		version := strings.Replace(tags[i], prefix, "", 1)
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
