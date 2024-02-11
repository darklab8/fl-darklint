package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../..")
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(stdout), "Available Commands") {
		t.Error("Help command did not find help info in output", string(stdout))
	}

	fmt.Println(string(stdout))
}
