package vergen

import (
	"fmt"
	"github.com/datawire/vergen/pkg/vergen"
	"github.com/spf13/cobra"
)

var branch string
var authority string
var appendRevision bool

func createPreviewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preview",
		Short: "Computes the preview version",
		Run:   createPreviewVersion,
	}

	cmd.Flags().StringVar(&branch, "branch", "", "The branch to use in the generated version")
	cmd.Flags().StringVar(&authority, "authority", "", "The user or system creating the version")
	cmd.Flags().BoolVar(&appendRevision, "revision", false, "Append an automatically generated revision ID to the preview version")

	return cmd
}

func createPreviewVersion(command *cobra.Command, args []string) {
	generator, err := vergen.NewGenerator()
	checkIfError(err)

	generator.ReleaseTagPrefix = prefix

	version, err := generator.PreviewVersion(branch, authority, appendRevision)
	checkIfError(err)

	fmt.Printf(version.String())
}
