package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fallback string

func LatestVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "latest",
		Short: "Examines Git tags to determine the latest release version",
		Run:   getCurrentVersion,
	}
}

func init() {
	cmd := LatestVersionCmd()

	RootCmd.AddCommand(cmd)

	cmd.Flags().StringVar(&fallback, "fallback", "", "Fallback value if there is no known \"latest\" version")
	cmd.Flags().StringVar(&prefix, "prefix", "release/", "A prefix the comes before the version")
}

func getCurrentVersion(cmd *cobra.Command, args []string) {
	repo := GetRepo()
	tags := GetTagNames(repo, prefix)

	result := getCurrentReleaseVersion(tags)

	if result == "" && fallback != "" {
		result = fallback
	}

	fmt.Println(result)
}
