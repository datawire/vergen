package vergen

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "vergen",
	Short: "Version Generator",
}

var prefix string

func init() {
	RootCmd.AddCommand(createPreviewCommand())
	RootCmd.AddCommand(createReleaseVersionCommand())
	RootCmd.AddCommand(createGetLatestVersion())
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		checkIfError(err)
	}
}

func checkIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
