package systems

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"strings"
)

type Config struct {
}

func (frelconfig *Config) Read(universe_config *universe.Config, filesystem filefind.Filesystem) *Config {

	var system_files map[string]*utils.File = make(map[string]*utils.File)
	for _, base := range universe_config.Bases {
		filename := universe_config.SystemMap[universe.SystemNickname(base.System)].File.FileName()
		path := filesystem.GetFile(strings.ToLower(filename))
		system_files[base.System] = &(utils.File{Filepath: path.Filepath})
	}

	var system_iniconfigs map[string]inireader.INIFile = make(map[string]inireader.INIFile)
	for system_key, file := range system_files {
		system := inireader.INIFile{}
		system_iniconfigs[system_key] = inireader.INIFile.Read(system, file)
	}

	return frelconfig

}

func (frelconfig *Config) Write() []*utils.File {
	var files []*utils.File = make([]*utils.File, 0)

	return files
}
