package market

import (
	"darktool/tools/utils"

	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := &utils.File{Filepath: filepath.Join(test_directory, FILENAME_SHIPS)}
	config := Config{}
	loaded_market_ships := config.Read(fileref)

	assert.Greater(t, len(loaded_market_ships.BaseGoods), 0, "market ships sections were not scanned")
	// TODO implement
	// assert.Greater(t, len(loaded_market_ships.BaseGoods[0].Goods), 0, "market ships sections were not scanned")
}

func TestWriter(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	input_file := &utils.File{Filepath: filepath.Join(test_directory, FILENAME_SHIPS)}

	temp_directory := utils.GetCurrrentTempFolder()

	config := Config{}
	config.Read(input_file)
	config.SetOutputPath(filepath.Join(temp_directory, FILENAME_SHIPS))
	output_file := config.Write()

	dry_run := true
	output_file.WriteLines(dry_run)
}
