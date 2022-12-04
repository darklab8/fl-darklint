package randline

import (
	"darktool/tools/utils"
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

type Input struct {
	InputFilePath  string
	OutputFilePath string
	Times          *int
}

func Run(input Input) {
	log.Info("Starting randLine")
	log.Info("InputFile=", input.InputFilePath)
	log.Info("OutputFile=", input.OutputFilePath)
	log.Info("Times=", input.Times)

	input_file := (&utils.File{Filepath: input.InputFilePath}).OpenToReadF()
	defer input_file.Close()
	input_lines := input_file.ReadLines()

	// write result

	output_file := (&utils.File{Filepath: input.OutputFilePath}).CreateToWriteF()
	defer output_file.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *input.Times; i++ {
		randomIndex := rand.Intn(len(input_lines))
		output_file.WritelnF(input_lines[randomIndex])
	}

	fmt.Println("OK")
}
