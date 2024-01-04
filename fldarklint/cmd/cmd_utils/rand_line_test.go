package cmd_utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"darklint/fldarklint/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/utils"
)

func TestRandLine(t *testing.T) {
	err := os.MkdirAll(filepath.Join(string(utils.GetCurrentFolder()), "rand_line", "tests", "temp"), 0777)
	logus.Log.CheckError(err, "failed to create folder")
	arguments := strings.Split("run . utils rand_line --input fldarklint/cmd/cmd_utils/rand_line/tests/data/input.txt --output fldarklint/cmd/cmd_utils/rand_line/tests/temp/output.txt -k 5", " ")
	cmd := exec.Command("go", arguments...)
	cmd.Dir = "../../.."
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	result := string(stdout)
	fmt.Println(result)
}
