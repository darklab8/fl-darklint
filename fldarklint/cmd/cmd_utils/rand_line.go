package cmd_utils

import (
	"github.com/spf13/cobra"

	randline "darklint/cmd/cmd_utils/rand_line"
)

// randLineCmd represents the randLine command
var randLineCmd = &cobra.Command{
	Use:   "rand_line",
	Short: "Tool to select `k` random lines from one file and copy to another one",
	Run: func(cmd *cobra.Command, args []string) {
		randline.Run(Input)
	},
}

var Input randline.Input

func init() {
	hookCmd.AddCommand(randLineCmd)

	randLineCmd.PersistentFlags().StringVarP(&Input.InputFilePath, "input", "i", "", "input input (required)")
	randLineCmd.MarkPersistentFlagRequired("input")

	randLineCmd.PersistentFlags().StringVarP(&Input.OutputFilePath, "output", "o", "", "output input (required)")
	randLineCmd.MarkPersistentFlagRequired("output")

	Input.Times = randLineCmd.PersistentFlags().IntP("k", "k", 0, "k-times elements to select randomly to new file (required)")
	randLineCmd.MarkPersistentFlagRequired("k")
}
