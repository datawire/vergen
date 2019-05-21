package vergen

import (
	"fmt"
	"github.com/plombardi89/vergen/pkg/vergen"
	"github.com/spf13/cobra"
)

func createReleaseVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next",
		Short: "Examines Git tags to determine the next release version",
		Run:   nextReleaseVersion,
	}

	cmd.Flags().StringVar(&prefix, "prefix", "release/", "A prefix the comes before the version")

	return cmd
}

func nextReleaseVersion(cmd *cobra.Command, args []string) {
	generator, err := vergen.NewGenerator()
	checkIfError(err)

	generator.ReleaseTagPrefix = prefix

	version, err := generator.NextReleaseVersion()
	checkIfError(err)

	fmt.Println(version.String())
}
