package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var version string
var buildDate string

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of license",
	Long:  `Print the version number of license`,
	Run: func(cmd *cobra.Command, args []string) {
		if version == "" {
			version = "development"
			buildDate = "just now"
		} else {
			version = "v" + version
		}
		fmt.Printf("license %s\nbuild date %s", version, buildDate)
	},
}
