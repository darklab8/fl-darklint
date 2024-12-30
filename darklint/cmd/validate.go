package cmd

import (
	"fmt"

	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-configs/configs/configs_export/trades"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darklint/darklint/common"
	"github.com/darklab8/fl-darklint/darklint/formatter"
	"github.com/darklab8/fl-darklint/darklint/validator"
	"github.com/darklab8/go-utils/utils/ptr"
	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate freelancer config files for being correct",
	Long:  `Freelancer folder is automatically discovered in any child folders`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate is called")

		configs := common.GetConfigs(GetFreelancerFolder())
		formatter.Run(configs, GetFreelancerFolder(), configs_mapped.IsDruRun(true))
		fmt.Println("formatter dry run did run")
		exported := configs_export.Export(configs, configs_export.ExportOptions{
			MappingOptions: trades.MappingOptions{SimplifiedTradeLanesCalc: ptr.Ptr(true)},
		})
		fmt.Println("export did run")
		validator.NewValidator(configs, exported).Run()
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.PersistentFlags().StringVarP(&freelancer_folder, "search", "s", freelancer_folder, "Freelancer location to search for validate running")
}
