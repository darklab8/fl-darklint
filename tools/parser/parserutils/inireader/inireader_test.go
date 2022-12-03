package inireader

import (
	"darktool/tools/utils"
	"path/filepath"
	"testing"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	marketships_filepath := filepath.Join(test_directory, "market_ships.ini")
	INIFileRead(marketships_filepath)
}
