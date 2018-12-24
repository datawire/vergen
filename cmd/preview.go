package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var commitID string
var branch string
var authority string
var appendRevision bool

func PreviewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "preview",
		Short: "Computes the preview version",
		Run:   createPreviewVersion,
	}
}

func createPreviewVersion(command *cobra.Command, args []string) {
	revision := ""
	if appendRevision {
		revision = "-" + generateRevision()
	}

	fmt.Printf("%s.%s.%s%s\n", commitID, branch, authority, revision)
}

func init() {
	cmd := PreviewCmd()

	RootCmd.AddCommand(cmd)

	cmd.Flags().StringVar(&commitID, "commit", "unknown", "Git commit ID")
	cmd.Flags().StringVar(&branch, "branch", "unknown", "Git branch name")
	cmd.Flags().StringVar(&authority, "authority", "unknown", "The user or system creating the version")
	cmd.Flags().BoolVar(&appendRevision, "revision", false, "Appends a generated revision ID to the preview version")
}

func generateRevision() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
