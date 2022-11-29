/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// randomizerCmd represents the randomizer command
var randomizerCmd = &cobra.Command{
	Use:   "randomizer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Starting randomizer")
		log.Info("InputFile=", InputFilePath)
		log.Info("OutputFile=", OutputFilePath)
		log.Info("Times=", Times)

		input_file, err := os.Open(InputFilePath)
		defer input_file.Close()
		if err != nil {
			log.Fatalf("failed to open")

		}
		scanner := bufio.NewScanner(input_file)
		var text []string

		for scanner.Scan() {
			text = append(text, scanner.Text())
		}

		// write result

		output_file, err := os.Create(OutputFilePath)

		if err != nil {
			log.Panic(err)
		}

		defer output_file.Close()

		n, err := strconv.Atoi(Times)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < n; i++ {
			randomIndex := rand.Intn(len(text))
			_, err2 := output_file.WriteString(fmt.Sprintf("%v\n", text[randomIndex]))

			if err2 != nil {
				log.Fatal(err2)
			}
		}
	},
}

var InputFilePath string
var OutputFilePath string
var Times string

func init() {
	rootCmd.AddCommand(randomizerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomizerCmd.PersistentFlags().StringVarP(&Foo, "foo", "", "A help for foo")

	randomizerCmd.PersistentFlags().StringVarP(&InputFilePath, "input", "i", "", "input input (required)")
	randomizerCmd.MarkPersistentFlagRequired("input")

	randomizerCmd.PersistentFlags().StringVarP(&OutputFilePath, "output", "o", "", "output input (required)")
	randomizerCmd.MarkPersistentFlagRequired("output")

	randomizerCmd.PersistentFlags().StringVarP(&Times, "k", "k", "", "k-times elements to select randomly to new file (required)")
	randomizerCmd.MarkPersistentFlagRequired("k")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomizerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
