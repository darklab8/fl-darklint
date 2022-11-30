/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"darktool/tools/randline"
)

// randLineCmd represents the randLine command
var randLineCmd = &cobra.Command{
	Use:   "randLine",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		randline.Run()
	},
}

func init() {
	rootCmd.AddCommand(randLineCmd)

	randLineCmd.PersistentFlags().StringVarP(&randline.InputFilePath, "input", "i", "", "input input (required)")
	randLineCmd.MarkPersistentFlagRequired("input")

	randLineCmd.PersistentFlags().StringVarP(&randline.OutputFilePath, "output", "o", "", "output input (required)")
	randLineCmd.MarkPersistentFlagRequired("output")

	randLineCmd.PersistentFlags().StringVarP(&randline.Times, "k", "k", "", "k-times elements to select randomly to new file (required)")
	randLineCmd.MarkPersistentFlagRequired("k")
}
