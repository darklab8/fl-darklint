/*
parse universe.ini
*/
package universe

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	FILENAME      = "universe.ini"
	KEY_BASETAG   = "[Base]"
	KEY_NICKNAME  = "nickname"
	KEY_STRIDNAME = "strid_name"
)

type Base struct {
	Nickname  inireader.ValueString
	StridName inireader.UniValue
}

type BaseNickname string

type Config struct {
	Bases []*Base

	BasesMap map[BaseNickname]*Base //key is
}

var Loaded *Config

func (frelconfig *Config) Read(input_file *utils.File) (*Config, inireader.INIFile) {
	if frelconfig.BasesMap == nil {
		frelconfig.BasesMap = make(map[BaseNickname]*Base)
	}

	if frelconfig.Bases == nil {
		frelconfig.Bases = make([]*Base, 0)
	}

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)

	bases, ok := iniconfig.SectionMap[KEY_BASETAG]
	if !ok {
		log.Trace("failed to access iniconfig.SectionMap", KEY_BASETAG)
	}
	for _, base := range bases {
		base_to_add := Base{}

		check_nickname := base.ParamMap[KEY_NICKNAME][0].First.(inireader.ValueString).AsString()
		if !utils.IsLower(check_nickname) {
			log.Warn("nickname: ", check_nickname, "in file universe.txt is not in lower case. Autofixing")
		}
		base_to_add.Nickname = base.ParamMap[KEY_NICKNAME][0].First.(inireader.ValueString).ToLowerValue()
		base_to_add.StridName = base.ParamMap[KEY_STRIDNAME][0].First

		frelconfig.Bases = append(frelconfig.Bases, &base_to_add)
		frelconfig.BasesMap[BaseNickname(base_to_add.Nickname)] = &base_to_add
	}

	return frelconfig, iniconfig
}

func Load() {
	file := &utils.File{Filepath: filefind.FreelancerFolder.Hashmap[FILENAME].Filepath}
	config := Config{}
	Loaded, _ = config.Read(file)
	log.Info(fmt.Sprintf("OK file.Filepath=%v, universe.ini is parsed to specialized data structs", file.Filepath))
}
