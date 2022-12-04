package inireader

import (
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := &utils.File{Filepath: filepath.Join(test_directory, "market_ships.ini")}
	config := INIFile.Read(INIFile{}, fileref)

	assert.Greater(t, len(config.Sections), 0, "market ships sections were not scanned")
}
