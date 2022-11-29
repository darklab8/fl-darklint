/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// randCoordsCmd represents the randCoords command
var randCoordsCmd = &cobra.Command{
	Use:   "randCoords",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("randCoords called")
		log.Info("Delimiter=", Delimiter)
		log.Info("RoundedNumbers=", RoundedNumbers)

		rand.Seed(time.Now().UnixNano())

		x := rand.Float64() * 180
		y := rand.Float64() * 180
		z := rand.Float64() * 180
		var sb strings.Builder
		roundedNeeded, _ := strconv.Atoi(RoundedNumbers)
		sb.WriteString(fmt.Sprintf("%v", toFixed(x, roundedNeeded)))
		sb.WriteString(fmt.Sprintf("%v", Delimiter))
		sb.WriteString(fmt.Sprintf("%v", toFixed(y, roundedNeeded)))
		sb.WriteString(fmt.Sprintf("%v", Delimiter))
		sb.WriteString(fmt.Sprintf("%v", toFixed(z, roundedNeeded)))

		fmt.Printf(fmt.Sprintf("%v\n", sb.String()))
	},
}

var Delimiter string
var RoundedNumbers string

func init() {
	rootCmd.AddCommand(randCoordsCmd)

	// set delimiter
	randCoordsCmd.PersistentFlags().StringVarP(&Delimiter, "delimiter", "d", ", ", "delimiter to separate")
	randCoordsCmd.PersistentFlags().StringVarP(&RoundedNumbers, "rounded_to", "r", "1", "rounded_numbers")
}
