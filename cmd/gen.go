package cmd

import "github.com/spf13/cobra"

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use: "gen",
}

func AddGeneratorCommand(parent *cobra.Command) {
	AddVersionCommand(genCmd)
	AddManifestCommand(genCmd)
	parent.AddCommand(genCmd)
}
