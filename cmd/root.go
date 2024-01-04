/*
package of binding all out scripts into CLI interface to user.

P.S. It should be containing as zero code logic as possible
*/
package cmd

import (
	"darklint/cmd/cmd_utils"
	"darklint/fldarklint/settings"
	"os"

	"github.com/spf13/cobra"
)

const description = "set of tools for config development of Freelancer game"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   settings.ToolName,
	Short: description,
	Long:  description + ``,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
//
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd_utils.Hook(rootCmd)
}
