package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "osl",
	Short: "osl is a cli tool to download license files",
	Long: `Download license files for your open source project.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
