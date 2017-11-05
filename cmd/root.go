package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "license [sub]",
	Short: "license is a cli tool to download license files",
	Long: `Download license files for your open source project.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			println("Usage: license [sub]")
			println("See `license help` for more")
		}
	},
}
