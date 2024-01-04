package cmd_utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use: "utils",
	Short: `utils extra commands useful during config development
		Contains nested sub commands!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK utils called")
		cmd.Help()
	},
}

func Hook(rootCmd *cobra.Command) {
	rootCmd.AddCommand(hookCmd)
}
