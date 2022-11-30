package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	cmd := exec.Command("go", "run", "../.")
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(stdout), "Available Commands") {
		t.Error("Help command did not find help info in output", string(stdout))
	}

	fmt.Println(string(stdout))
}

func TestRandRotator(t *testing.T) {
	cmd := exec.Command("go", "run", "../.", "randRotator")
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(stdout), ", ") {
		t.Error("string does not contain resired substring", string(stdout))
	}

	fmt.Println(string(stdout))
}

func TestRandLine(t *testing.T) {
	arguments := strings.Split("run . randLine --input tools/randline/tests/data/input.txt --output tools/randline/tests/temp/output.txt -k 5", " ")
	cmd := exec.Command("go", arguments...)
	cmd.Dir = ".."
	stdout, err := cmd.Output()

	if err != nil {
		t.Error(err)
	}

	result := string(stdout)
	fmt.Println(result)
}
