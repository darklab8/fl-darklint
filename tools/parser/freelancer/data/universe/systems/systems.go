package systems

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"strings"
)

const (
	KEY_OBJECT   = "[Object]"
	KEY_NICKNAME = "nickname"
	KEY_BASE     = "base"
)

type Base struct {
	Nickname string
	Base     string // base.nickname in universe.ini
}
type System struct {
	Nickname    string
	Bases       []*Base
	BasesByNick map[string]*Base
	BasesByBase map[string]*Base
}

type Config struct {
	SystemsMap map[string]*System
	Systems    []*System
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

	frelconfig.SystemsMap = make(map[string]*System)
	frelconfig.Systems = make([]*System, 0)
	for system_key, sysiniconf := range system_iniconfigs {
		system_to_add := System{}
		system_to_add.Nickname = system_key
		system_to_add.BasesByNick = make(map[string]*Base)
		system_to_add.BasesByBase = make(map[string]*Base)
		system_to_add.Bases = make([]*Base, 0)
		frelconfig.SystemsMap[system_key] = &system_to_add
		frelconfig.Systems = append(frelconfig.Systems, &system_to_add)

		objects, ok := sysiniconf.SectionMap[KEY_OBJECT]
		if ok {
			for _, obj := range objects {

				// check if it is base object
				_, ok := obj.ParamMap[KEY_BASE]
				if ok {
					base_to_add := Base{}

					base_to_add.Nickname = obj.GetParamStrToLower(KEY_NICKNAME, inireader.REQUIRED_p)
					base_to_add.Base = obj.GetParamStrToLower(KEY_BASE, inireader.REQUIRED_p)
					system_to_add.BasesByBase[base_to_add.Base] = &base_to_add
					system_to_add.BasesByNick[base_to_add.Nickname] = &base_to_add
					system_to_add.Bases = append(system_to_add.Bases, &base_to_add)
				}
			}
		}

	}

	return frelconfig
}

func (frelconfig *Config) Write() []*utils.File {
	var files []*utils.File = make([]*utils.File, 0)

	return files
}
