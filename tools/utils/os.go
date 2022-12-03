package utils

import (
	"path/filepath"
	"runtime"
)

// example
func GetCurrrentTestFolder() string {
	_, filename, _, _ := runtime.Caller(1)
	directory := filepath.Dir(filename)
	test_directory := filepath.Join(directory, "testdata")
	return test_directory
}
