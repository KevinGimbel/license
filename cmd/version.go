package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/kevingimbel/license/lib"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of license",
	Long:  `Print the version number of license`,
	Run: func(cmd *cobra.Command, args []string) {
		var version = lib.GetVersion()
		var buildDate = lib.GetBuildDate()

		if version == "" {
			version = "development"
			buildDate = "just now"
		} else {
			version = "v" + version
		}
		fmt.Printf("license %s\nbuild date %s", version, buildDate)
	},
}
