package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/kevingimbel/license/lib"
)

var (
	long bool
)

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&long, "long", "l",false, "Display full version info")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of license",
	Long:  `Print the version of license. Use "-l" to display build date and git commit info.`,
	Run: func(cmd *cobra.Command, args []string) {
		var version = lib.GetVersion()
		var buildDate = lib.GetBuildDate()
		var commit = lib.GetCommit()

		// If version is not set we run in dev mode / compile locally
		if version == "" {
			version = "development"
			buildDate = "just now"
			commit = ""
		}

		if long {
			fmt.Printf("Version: %s\nCommit: %s\nBuild Date: %s", version, commit, buildDate)
		} else {
			fmt.Println(version)
		}
	},
}
