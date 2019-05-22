package vergen

import (
	"fmt"
	"github.com/datawire/vergen/pkg/vergen"
	"github.com/spf13/cobra"
)

var fallback string

func createGetLatestVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest",
		Short: "Examines Git tags to determine the latest release version",
		Run:   getCurrentVersion,
	}

	cmd.Flags().StringVar(&fallback, "fallback", "", "Fallback value if there is no known \"latest\" version")
	cmd.Flags().StringVar(&prefix, "prefix", "release/", "A prefix the comes before the version")

	return cmd
}

func getCurrentVersion(cmd *cobra.Command, args []string) {
	generator, err := vergen.NewGenerator()
	checkIfError(err)

	generator.ReleaseTagPrefix = prefix

	current, err := generator.LatestReleaseVersion(fallback)
	checkIfError(err)

	fmt.Printf(current)
}
