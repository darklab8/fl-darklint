/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/darklab8/fl-darklint/darklint/settings"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "see current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("darklint version: v%s\n", settings.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
