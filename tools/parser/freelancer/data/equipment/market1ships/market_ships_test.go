package market1ships

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
	fileref := &utils.File{Filepath: filepath.Join(test_directory, Filename)}
	config := Config{}
	loaded_market_ships := config.Read(fileref)

	assert.Greater(t, len(loaded_market_ships.BaseGoods), 0, "market ships sections were not scanned")
	assert.Greater(t, len(loaded_market_ships.BaseGoods[0].Goods), 0, "market ships sections were not scanned")
}

func TestWriter(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	input_file := &utils.File{Filepath: filepath.Join(test_directory, Filename)}

	temp_directory := utils.GetCurrrentTempFolder()
	output_file := &utils.File{Filepath: filepath.Join(temp_directory, Filename)}

	config := Config{}
	config.Read(input_file)
	config.Write(output_file)

	output_file.WriteLines()
}

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(folder)))
	log.Debug(freelancer_folder)
	filesystem := filefind.FindConfigs(freelancer_folder)

	temp_directory := utils.GetCurrrentTempFolder()
	output_config := &utils.File{Filepath: filepath.Join(temp_directory, Filename)}

	market_config := Config{}
	market_config.Read(&utils.File{Filepath: filesystem.Hashmap[Filename].Filepath})

	universe_config := universe.Config{}
	universe_config.Read(&utils.File{Filepath: filesystem.Hashmap[universe.Filename].Filepath})

	info_config := service.Config{}
	info_config.Read(&utils.File{Filepath: filesystem.Hashmap[service.Filename].Filepath})

	market_config.UpdateWithBasenames(&universe_config, &info_config)
	market_config.Write(output_config)

	// isRecycleCandidate
	isRecyclePresent := false
	lines := output_config.GetLines()
	for _, line := range lines {
		if strings.Contains(line, "isRecycleCandidate") {
			isRecyclePresent = true
		}
	}

	output_config.WriteLines()

	assert.True(t, isRecyclePresent)
}
