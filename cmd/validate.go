/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"darktool/tools/validator"
	"fmt"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate freelancer config files for being correct and if not, try to fix",
	Long: `Freelancer folder is automatically discovered in any child folders
or you can set its location with ENV variable DARKTOOL_FREELANCER_FOLDER`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate called")
		validator.Run()
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
