/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"darklint/fldarklint/settings"
	"darklint/fldarklint/validator"
	"fmt"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/configs_mapped"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "format",
	Short: "format freelancer config files for being correct",
	Long:  `Freelancer folder is automatically discovered in any child folders`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("format called")
		validator.Run(configs_mapped.IsDruRun(is_dry_run))
		fmt.Println("OK")
	},
}

var is_dry_run bool

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.PersistentFlags().BoolVarP(&is_dry_run, "dry", "d", false, "enable dry for checks without writing to file / good for CI")
	validateCmd.PersistentFlags().StringVarP(&settings.FreelancerFreelancerLocation, "search", "s", settings.FreelancerFreelancerLocation, "Freelancer location to search for validate running")
}
