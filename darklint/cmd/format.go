/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/darklab8/fl-darklint/darklint/common"
	"github.com/darklab8/fl-darklint/darklint/formatter"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "format freelancer config files for being correct",
	Long:  `Freelancer folder is automatically discovered in any child folders`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("format called")

		configs := common.GetConfigs(GetFreelancerFolder())
		formatter.Run(configs, GetFreelancerFolder(), configs_mapped.IsDruRun(is_dry_run))
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.PersistentFlags().BoolVarP(&is_dry_run, "dry", "d", false, "enable dry for checks without writing to file / good for CI")
	formatCmd.PersistentFlags().StringVarP(&freelancer_folder, "search", "s", freelancer_folder, "Freelancer location to search for validate running")
}
