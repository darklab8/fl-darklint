/*
package of binding all out scripts into CLI interface to user.

P.S. It should be containing as zero code logic as possible
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const description = "set of tools for config development of Freelancer game"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "darktool",
	Short: description,
	Long: description + `

download latest release at
https://github.com/darklab8/darklab_freelancer_darktool/releases
check 'darktool version' for which one u have at the moment

'darktool validate' should be executed from root of your Freelancer folder
it will parse and validate different files, and automatically autofix them.
run with 'darktool validate --dry' in order to be not overwriting files
P.S. Freelancer folder location can be overwritten with env flag DARKTOOL_FREELANCER_FOLDER`,

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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.darktool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
