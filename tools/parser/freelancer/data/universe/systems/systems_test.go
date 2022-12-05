package systems

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(folder))))
	log.Debug(freelancer_folder)
	filesystem := filefind.FindConfigs(freelancer_folder)

	universe_config := universe.Config{}
	universe_config.Read(&utils.File{Filepath: filesystem.Hashmap[universe.FILENAME].Filepath})

	systems := (&Config{}).Read(&universe_config, filesystem)
	_ = systems
}
