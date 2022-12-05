package systems

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"strings"
)

type Config struct {
}

func (frelconfig *Config) Read(universe_config *universe.Config, filesystem filefind.Filesystem) *Config {

	var input_systems []*utils.File = make([]*utils.File, 0)
	for _, base := range universe_config.Bases {
		filename := universe_config.SystemMap[universe.SystemNickname(base.System)].File.FileName()
		path := filesystem.GetFile(strings.ToLower(filename))
		input_systems = append(input_systems, &(utils.File{Filepath: path.Filepath}))
	}

	return frelconfig

}

func (frelconfig *Config) Write() []*utils.File {
	var files []*utils.File = make([]*utils.File, 0)

	return files
}
