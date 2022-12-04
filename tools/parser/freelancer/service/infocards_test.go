package service

import (
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := &utils.File{Filepath: filepath.Join(test_directory, FILENAME)}
	config := Config{}
	config.Read(fileref)

	assert.Greater(t, len(config.Records), 0)
}

func TestInterfaceStruct(t *testing.T) {
	var record Record

	record = Infocard{}

	record = Name{}
	assert.True(t, record != nil)
}
