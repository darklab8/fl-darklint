package randline

import (
	"testing"
)

func TestSimple(t *testing.T) {
	input := Input{
		InputFilePath:  "tests/data/input.txt",
		OutputFilePath: "tests/temp/output.txt",
		Times:          "5",
	}
	Run(input)
}
