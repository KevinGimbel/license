package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "license",
	Short: "license is a cli tool to download license files",
	Long: `Download license files for your open source project by `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
