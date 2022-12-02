package cmd

import (
	"darktool/tools/randrotator"
	"fmt"

	"github.com/spf13/cobra"
)

// randRotatorCmd represents the randRotator command
var randRotatorCmd = &cobra.Command{
	Use:   "randRotator",
	Short: "Generate random rotation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(randrotator.Run(inputCMD))
	},
}

var inputCMD randrotator.Input

func init() {
	rootCmd.AddCommand(randRotatorCmd)

	randRotatorCmd.PersistentFlags().StringVarP(&inputCMD.Delimiter, "delimiter", "d", ", ", "delimiter to separate")
	inputCMD.RoundedNumbers = randRotatorCmd.PersistentFlags().IntP("rounded_to", "r", 1, "rounded_numbers")
}
