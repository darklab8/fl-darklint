package utils

import (
	"path/filepath"
	"runtime"
)

func GetCurrrentChildFolder(folder_name string) string {
	_, filename, _, _ := runtime.Caller(2)
	directory := filepath.Dir(filename)
	test_directory := filepath.Join(directory, folder_name)
	return test_directory
}

func GetCurrrentTestFolder() string {
	return GetCurrrentChildFolder("testdata")
}

func GetCurrrentTempFolder() string {
	return GetCurrrentChildFolder("tempdata")
}

func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	directory := filepath.Dir(filename)
	test_directory := filepath.Join(directory, "testdata")
	return test_directory
}
