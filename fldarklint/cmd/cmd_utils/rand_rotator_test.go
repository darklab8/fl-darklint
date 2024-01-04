package cmd_utils

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestRandRotator(t *testing.T) {
	cmd := exec.Command("go", "run", "../../..", "utils", "rand_rotator")
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(stdout), ", ") {
		t.Error("string does not contain resired substring", string(stdout))
	}

	fmt.Println(string(stdout))
}
