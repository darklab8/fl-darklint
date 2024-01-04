package rand_line

import (
	"testing"
)

func TestSimple(t *testing.T) {
	times := 5
	input := Input{
		InputFilePath:  "tests/data/input.txt",
		OutputFilePath: "tests/temp/output.txt",
		Times:          &times,
	}
	Run(input)
}
