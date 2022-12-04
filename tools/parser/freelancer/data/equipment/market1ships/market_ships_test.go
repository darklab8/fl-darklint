package market1ships

import (
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := &utils.File{Filepath: filepath.Join(test_directory, filename)}
	config := Config{}
	loaded_market_ships := config.Read(fileref)

	assert.Greater(t, len(loaded_market_ships.BaseGoods), 0, "market ships sections were not scanned")
	assert.Greater(t, len(loaded_market_ships.BaseGoods[0].Goods), 0, "market ships sections were not scanned")
}

func TestWriter(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	input_file := &utils.File{Filepath: filepath.Join(test_directory, filename)}

	temp_directory := utils.GetCurrrentTempFolder()
	output_file := &utils.File{Filepath: filepath.Join(temp_directory, filename)}

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	_ = iniconfig

	config := Config{}
	config.Read(input_file)
	config.Write(output_file)

	output_file.WriteLines()
}
