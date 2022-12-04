package market

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"strings"

	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := &utils.File{Filepath: filepath.Join(test_directory, FILENAME_SHIPS)}
	config := Config{}
	loaded_market_ships := config.Read(fileref)

	assert.Greater(t, len(loaded_market_ships.BaseGoods), 0, "market ships sections were not scanned")
	assert.Greater(t, len(loaded_market_ships.BaseGoods[0].Goods), 0, "market ships sections were not scanned")
}

func TestWriter(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	input_file := &utils.File{Filepath: filepath.Join(test_directory, FILENAME_SHIPS)}

	temp_directory := utils.GetCurrrentTempFolder()
	output_file := &utils.File{Filepath: filepath.Join(temp_directory, FILENAME_SHIPS)}

	config := Config{}
	config.Read(input_file)
	config.Write(output_file)

	dry_run := true
	output_file.WriteLines(dry_run)
}

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(folder)))
	log.Debug(freelancer_folder)
	filesystem := filefind.FindConfigs(freelancer_folder)

	temp_directory := utils.GetCurrrentTempFolder()
	output_config := &utils.File{Filepath: filepath.Join(temp_directory, FILENAME_SHIPS)}

	market_config := Config{}
	market_config.Read(&utils.File{Filepath: filesystem.Hashmap[FILENAME_SHIPS].Filepath})

	universe_config := universe.Config{}
	universe_config.Read(&utils.File{Filepath: filesystem.Hashmap[universe.FILENAME].Filepath})

	info_config := service.Config{}
	info_config.Read(&utils.File{Filepath: filesystem.Hashmap[service.FILENAME].Filepath})

	market_config.UpdateWithBasenames(&universe_config, &info_config)
	market_config.Write(output_config)

	// isRecycleCandidate
	isRecyclePresent := false
	lines := output_config.GetLines()
	for _, line := range lines {
		if strings.Contains(line, KEY_RECYCLE) {
			isRecyclePresent = true
		}
	}

	dry_run := true
	output_config.WriteLines(dry_run)

	assert.True(t, isRecyclePresent)
}
