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
	Nickname  inireader.ValueString
	StridName inireader.UniValue
}

type BaseNickname string

type Config struct {
	Bases []*Base

	BasesMap map[BaseNickname]*Base //key is
}

var Loaded Config

func Read(input_file utils.File) Config {
	var frelconfig Config
	Loaded.BasesMap = make(map[BaseNickname]*Base)

	iniconfig := inireader.INIFileRead(input_file)

	bases, ok := iniconfig.SectionMap[BaseTag]
	if !ok {
		log.Trace("failed to access iniconfig.SectionMap[BaseTag]")
	}
	for _, base := range bases {
		base_to_add := Base{}

		check_nickname := base.ParamMap["nickname"][0].First.(inireader.ValueString).AsString()
		if !utils.IsLower(check_nickname) {
			log.Warn("nickname: ", check_nickname, "in file universe.txt is not in lower case. Autofixing")
		}
		base_to_add.Nickname = base.ParamMap["nickname"][0].First.(inireader.ValueString).ToLowerValue()
		base_to_add.StridName = base.ParamMap["strid_name"][0].First

		Loaded.Bases = append(Loaded.Bases, &base_to_add)
		Loaded.BasesMap[BaseNickname(base_to_add.Nickname)] = &base_to_add
	}

	return frelconfig
}

func Load() {
	file := utils.File{Filepath: filefind.FreelancerFolder.Hashmap[filename].Filepath}
	Read(file)
	log.Info("OK universe.ini is parsed to specialized data structs")
}
