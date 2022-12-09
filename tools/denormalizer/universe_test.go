package denormalizer

import (
	"darktool/tools/parser/freelancer/data/equipment/market"
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

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(folder)))
	log.Debug(freelancer_folder)
	filesystem := filefind.FindConfigs(freelancer_folder)

	temp_directory := utils.GetCurrrentTempFolder()

	market_config := market.Config{}
	market_config.Read(&utils.File{Filepath: filesystem.Hashmap[market.FILENAME_SHIPS].Filepath})

	universe_config := universe.Config{}
	universe_config.Read(&utils.File{Filepath: filesystem.Hashmap[universe.FILENAME].Filepath})

	info_config := infocard.Config{}
	info_config.Read(&utils.File{Filepath: filesystem.Hashmap[infocard.FILENAME].Filepath})

	systems_config := systems.Config{}
	systems_config.Read(&universe_config, filesystem)

	denormalizer := (&Denormalizer{}).Create(&universe_config)
	denormalizer.ReadBaseNames(&market_config, &universe_config, &info_config)
	denormalizer.ReadRecycle(&market_config, &universe_config, &systems_config)
	denormalizer.MarketWrite(&market_config)

	market_config.SetOutputPath(filepath.Join(temp_directory, market.FILENAME_SHIPS))
	output_config := market_config.Write()

	// isRecycleCandidate
	isRecyclePresent := false
	lines := output_config.GetLines()
	for _, line := range lines {
		if strings.Contains(line, market.KEY_RECYCLE) {
			isRecyclePresent = true
		}
	}

	dry_run := true
	output_config.WriteLines(dry_run)

	assert.True(t, isRecyclePresent)
}
