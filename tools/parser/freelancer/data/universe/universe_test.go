/*
parse universe.ini
*/
package universe

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := &utils.File{Filepath: filepath.Join(test_directory, FILENAME)}
	config := Config{}
	config.Read(fileref)

	assert.Greater(t, len(config.Bases), 0)
	assert.Greater(t, len(config.Systems), 0)
}

func TestIdentifySystemFiles(t *testing.T) {
	test_directory := utils.GetCurrentFolder()
	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(test_directory)))
	filesystem := filefind.FindConfigs(freelancer_folder)
	log.Debug(filesystem)

	config := Config{}
	universe_fileref := &utils.File{Filepath: filepath.Join(test_directory, "testdata", FILENAME)}
	config.Read(universe_fileref)
	log.Debug("breakpoint")
}
