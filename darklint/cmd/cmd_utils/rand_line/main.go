package rand_line

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/darklab8/fl-darklint/darklint/settings/logus"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"

	"github.com/darklab8/go-utils/utils/utils_types"
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

	input_lines, err := file.NewFile(utils_types.FilePath(input.InputFilePath)).ReadLines()
	logus.Log.CheckPanic(err, "failed to read lines of a file")

	// write result

	output_file := file.NewFile(utils_types.FilePath(input.OutputFilePath))

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *input.Times; i++ {
		randomIndex := rand.Intn(len(input_lines))
		output_file.ScheduleToWrite(input_lines[randomIndex])
	}

	output_file.WriteLines()

	fmt.Println("OK")
}
