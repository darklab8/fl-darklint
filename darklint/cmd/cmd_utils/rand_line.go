package cmd_utils

import (
	"github.com/spf13/cobra"

	"github.com/darklab8/fl-darklint/darklint/cmd/cmd_utils/rand_line"
)

// randLineCmd represents the randLine command
var randLineCmd = &cobra.Command{
	Use:   "rand_line",
	Short: "Tool to select `k` random lines from one file and copy to another one",
	Run: func(cmd *cobra.Command, args []string) {
		rand_line.Run(Input)
	},
}

var Input rand_line.Input

func init() {
	hookCmd.AddCommand(randLineCmd)

	randLineCmd.PersistentFlags().StringVarP(&Input.InputFilePath, "input", "i", "", "input input (required)")
	randLineCmd.MarkPersistentFlagRequired("input")

	randLineCmd.PersistentFlags().StringVarP(&Input.OutputFilePath, "output", "o", "", "output input (required)")
	randLineCmd.MarkPersistentFlagRequired("output")

	Input.Times = randLineCmd.PersistentFlags().IntP("k", "k", 0, "k-times elements to select randomly to new file (required)")
	randLineCmd.MarkPersistentFlagRequired("k")
}
