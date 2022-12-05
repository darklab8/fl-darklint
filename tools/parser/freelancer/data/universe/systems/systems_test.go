package systems

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(folder))))
	log.Debug(freelancer_folder)
	filesystem := filefind.FindConfigs(freelancer_folder)

	universe_config := universe.Config{}
	universe_config.Read(&utils.File{Filepath: filesystem.Hashmap[universe.FILENAME].Filepath})

	systems := (&Config{}).Read(&universe_config, filesystem)

	system, ok := systems.SystemsMap["br01"]
	assert.True(t, ok, "system should be present")

	_, ok = system.BasesByBase["br01_01_base"]
	assert.True(t, ok, "base should be present")
}
