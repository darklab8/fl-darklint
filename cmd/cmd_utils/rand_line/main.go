package rand_line

import (
	"darklint/fldarklint/logus"
	"darklint/tools/parser/parserutils/filefind/file"
	"fmt"
	"math/rand"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

type Input struct {
	InputFilePath  string
	OutputFilePath string
	Times          *int
}

func Run(input Input) {
	logus.Log.Info("Starting randLine")
	logus.Log.Info("InputFile=" + input.InputFilePath)
	logus.Log.Info("OutputFile=" + input.OutputFilePath)
	logus.Log.Info(fmt.Sprintf("Times=%d", *input.Times))

	input_file := file.NewFile(utils_types.FilePath(input.InputFilePath)).OpenToReadF()
	defer input_file.Close()
	input_lines := input_file.ReadLines()

	// write result

	output_file := file.NewFile(utils_types.FilePath(input.OutputFilePath)).CreateToWriteF()
	defer output_file.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *input.Times; i++ {
		randomIndex := rand.Intn(len(input_lines))
		output_file.WritelnF(input_lines[randomIndex])
	}

	fmt.Println("OK")
}
