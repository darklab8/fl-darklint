/*
parse universe.ini
*/
package universe

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"

	log "github.com/sirupsen/logrus"
)

const (
	filename = "universe.ini"
	BaseTag  = "[Base]"
)

type Base struct {
	nickname   inireader.ValueString
	strid_name inireader.UniValue
}

type Config struct {
	Bases []Base
}

var LoadedConfig Config

func Read(input_file utils.File) Config {
	var frelconfig Config

	iniconfig := inireader.INIFileRead(input_file)

	bases, ok := iniconfig.SectionMap[BaseTag]
	if !ok {
		log.Trace("failed to access iniconfig.SectionMap[BaseTag]")
	}
	for _, base := range bases {
		base_to_add := Base{}
		base_to_add.nickname = base.ParamMap["nickname"][0].First.(inireader.ValueString)
		base_to_add.strid_name = base.ParamMap["strid_name"][0].First
	}

	return frelconfig
}

func Load() {
	file := utils.File{Filepath: filefind.FreelancerFolder.Hashmap[filename].Filepath}
	Read(file)
	log.Info("OK universe.ini is parsed to specialized data structs")
}
