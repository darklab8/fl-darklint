/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd_utils

import (
	"fmt"

	"darklint/cmd/cmd_utils/findduplicates"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/spf13/cobra"
)

// findDuplicatesCmd represents the findDuplicates command
var findDuplicatesCmd = &cobra.Command{
	Use:   "findDuplicates",
	Short: "Find duplicates by regular expression",
	Long: `
	Finds all matching occurencies by fuzzy regular expression
	Count them and show which ones exact matches are encountered more than once
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("findDuplicates called")
		fmt.Printf("filepath=%s\n", pathTofile)
		fmt.Printf("regex=%s\n", regexExpression)
		findduplicates.Main(utils_types.FilePath(pathTofile), regexExpression)
	},
}

var pathTofile string
var regexExpression string

func inif() {
	hookCmd.AddCommand(findDuplicatesCmd)
	findDuplicatesCmd.Flags().StringVarP(&pathTofile, "filepath", "f", "", "Path to file where duplicates to find")
	findDuplicatesCmd.Flags().StringVarP(&regexExpression, "regexp", "r", "", "Regular rexpression to match")
	findDuplicatesCmd.MarkFlagRequired("filepath")
	findDuplicatesCmd.MarkFlagRequired("regexp")
}
