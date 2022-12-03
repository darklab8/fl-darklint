/*
parse universe.ini
*/
package universe

import (
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := utils.File{Filepath: filepath.Join(test_directory, filename)}
	_ = Read(fileref)

	assert.Greater(t, len(LoadedConfig.Bases), 0)
}
