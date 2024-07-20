package rand_line

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/darklab8/go-utils/utils/utils_os"
)

func TestSimple(t *testing.T) {
	os.MkdirAll(filepath.Join(string(utils_os.GetCurrentFolder()), "tests", "temp"), 0777)
	times := 5
	input := Input{
		InputFilePath:  "tests/data/input.txt",
		OutputFilePath: "tests/temp/output.txt",
		Times:          &times,
	}
	Run(input)
}
