package cmd_utils

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestRandLine(t *testing.T) {
	arguments := strings.Split("run . utils rand_line --input cmd/cmd_utils/rand_line/tests/data/input.txt --output cmd/cmd_utils/rand_line/tests/temp/output.txt -k 5", " ")
	cmd := exec.Command("go", arguments...)
	cmd.Dir = "../../.."
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	result := string(stdout)
	fmt.Println(result)
}
