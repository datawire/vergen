package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var commitID string
var branch string
var authority string

func PreviewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "preview",
		Short: "Computes the preview version",
		Run:   createPreviewVersion,
	}
}

func createPreviewVersion(command *cobra.Command, args []string) {
	fmt.Printf("%s.%s.%s\n", commitID, branch, authority)
}

func init() {
	cmd := PreviewCmd()

	RootCmd.AddCommand(cmd)

	cmd.Flags().StringVar(&commitID, "commit", "unknown", "Git commit ID")
	cmd.Flags().StringVar(&branch, "branch", "unknown", "Git branch name")
	cmd.Flags().StringVar(&authority, "authority", "unknown", "The user or system creating the version")
}
