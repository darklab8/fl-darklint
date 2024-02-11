package cmd_utils

import (
	"fmt"

	"github.com/darklab8/fl-darklint/darklint/cmd/cmd_utils/rand_rotator"

	"github.com/spf13/cobra"
)

// randRotatorCmd represents the randRotator command
var randRotatorCmd = &cobra.Command{
	Use:   "rand_rotator",
	Short: "Generate random rotation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(rand_rotator.Run(inputCMD))
	},
}

var inputCMD rand_rotator.Input

func init() {
	hookCmd.AddCommand(randRotatorCmd)

	randRotatorCmd.PersistentFlags().StringVarP(&inputCMD.Delimiter, "delimiter", "d", ", ", "delimiter to separate")
	inputCMD.RoundedNumbers = randRotatorCmd.PersistentFlags().IntP("rounded_to", "r", 1, "rounded_numbers")
}
