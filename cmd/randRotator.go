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
		fmt.Printf(randrotator.Run(inputCMD))
	},
}

var inputCMD randrotator.Input

func init() {
	rootCmd.AddCommand(randRotatorCmd)

	randRotatorCmd.PersistentFlags().StringVarP(&inputCMD.Delimiter, "delimiter", "d", ", ", "delimiter to separate")
	inputCMD.RoundedNumbers = randRotatorCmd.PersistentFlags().IntP("rounded_to", "r", 1, "rounded_numbers")
}
