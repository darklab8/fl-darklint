/*
parse universe.ini
*/
package universe

import (
	"darktool/tools/utils"
	"path/filepath"
	"testing"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := utils.File{Filepath: filepath.Join(test_directory, filename)}
	_ = Read(fileref)

	// assert.Greater(t, len(parsed_data.Base_goods), 0)
	// assert.Greater(t, len(parsed_data.Base_goods[0].Goods), 0)
}
