/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"darktool/tools/randrotator"
	"fmt"

	"github.com/spf13/cobra"
)

// randRotatorCmd represents the randRotator command
var randRotatorCmd = &cobra.Command{
	Use:   "randRotator",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(randrotator.Run())
	},
}

func init() {
	rootCmd.AddCommand(randRotatorCmd)

	// set delimiter
	randRotatorCmd.PersistentFlags().StringVarP(&randrotator.Delimiter, "delimiter", "d", ", ", "delimiter to separate")
	randRotatorCmd.PersistentFlags().StringVarP(&randrotator.RoundedNumbers, "rounded_to", "r", "1", "rounded_numbers")
}
