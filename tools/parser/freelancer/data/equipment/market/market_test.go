package market

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/data/universe/systems"
	"darktool/tools/parser/freelancer/infocard"
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

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(folder)))
	log.Debug(freelancer_folder)
	filesystem := filefind.FindConfigs(freelancer_folder)

	temp_directory := utils.GetCurrrentTempFolder()

	market_config := Config{}
	market_config.Read(&utils.File{Filepath: filesystem.Hashmap[FILENAME_SHIPS].Filepath})

	universe_config := universe.Config{}
	universe_config.Read(&utils.File{Filepath: filesystem.Hashmap[universe.FILENAME].Filepath})

	info_config := infocard.Config{}
	info_config.Read(&utils.File{Filepath: filesystem.Hashmap[infocard.FILENAME].Filepath})

	systems_config := systems.Config{}
	systems_config.Read(&universe_config, filesystem)

	denormalizer := (&Denormalizer{}).Create(&universe_config)
	denormalizer.ReadBaseNames(&market_config, &universe_config, &info_config)
	denormalizer.ReadRecycle(&market_config, &universe_config, &systems_config)
	denormalizer.Write(&market_config)

	market_config.SetOutputPath(filepath.Join(temp_directory, FILENAME_SHIPS))
	output_config := market_config.Write()

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
