/*
parse universe.ini
*/
package universe

import (
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	FILENAME      = "universe.ini"
	KEY_BASE_TAG  = "[Base]"
	KEY_NICKNAME  = "nickname"
	KEY_STRIDNAME = "strid_name"
	KEY_SYSTEM    = "system"
	KEY_FILE      = "file"

	KEY_SYSTEM_TAG = "[system]"
)

// Linux friendly filepath, that can be returned to Windows way from linux
type Path string

func PathCreate(input string) Path {
	input = strings.ReplaceAll(input, `\`, `/`)
	input = strings.ToLower(input)
	return Path(input)
}

func (p Path) LinuxPath() string {
	return string(p)
}

func (p Path) WindowsPath() string {
	return strings.ReplaceAll(string(p), `/`, `\`)
}

type Base struct {
	Nickname  string
	System    string
	StridName int
	// Danger. filepath in Windows File path system
	// Hopefully filepath will read it
	File             Path
	BGCS_base_run_by string
}

type BaseNickname string

type SystemNickname string

type System struct {
	Nickname      string
	Pos           [2]int
	Msg_id_prefix string
	Visit         int
	Strid_name    int
	Ids_info      int
	File          Path
	NavMapScale   int
}

type Config struct {
	Bases    []*Base
	BasesMap map[BaseNickname]*Base //key is

	Systems   []*System
	SystemMap map[SystemNickname]*System //key is
}

func (frelconfig *Config) AddBase(base_to_add *Base) {
	frelconfig.Bases = append(frelconfig.Bases, base_to_add)
	frelconfig.BasesMap[BaseNickname(base_to_add.Nickname)] = base_to_add
}

func (frelconfig *Config) AddSystem(system_to_add *System) {
	frelconfig.Systems = append(frelconfig.Systems, system_to_add)
	frelconfig.SystemMap[SystemNickname(system_to_add.Nickname)] = system_to_add
}

func (frelconfig *Config) Read(input_file *utils.File) (*Config, inireader.INIFile) {
	if frelconfig.BasesMap == nil {
		frelconfig.BasesMap = make(map[BaseNickname]*Base)
	}

	if frelconfig.Bases == nil {
		frelconfig.Bases = make([]*Base, 0)
	}

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)

	bases, ok := iniconfig.SectionMap[KEY_BASE_TAG]
	if !ok {
		log.Trace("failed to access iniconfig.SectionMap", KEY_BASE_TAG)
	}
	for _, base := range bases {
		base_to_add := Base{}

		check_nickname := base.ParamMap[KEY_NICKNAME][0].First.(inireader.ValueString).AsString()
		if !utils.IsLower(check_nickname) {
			log.Warn("nickname: ", check_nickname, "in file universe.txt is not in lower case. Autofixing")
		}
		base_to_add.Nickname = strings.ToLower(base.ParamMap[KEY_NICKNAME][0].First.AsString())
		strid, err := strconv.Atoi(base.ParamMap[KEY_STRIDNAME][0].First.AsString())
		if err != nil {
			log.Fatal("failed to parse strid in universe.ini for base=", base_to_add)
		}
		base_to_add.StridName = strid

		base_to_add.System = strings.ToLower(base.ParamMap[KEY_SYSTEM][0].First.AsString())
		base_to_add.File = PathCreate(base.ParamMap[KEY_FILE][0].First.AsString())

		frelconfig.AddBase(&base_to_add)
	}

	// Systems
	if frelconfig.SystemMap == nil {
		frelconfig.SystemMap = make(map[SystemNickname]*System)
	}

	if frelconfig.Systems == nil {
		frelconfig.Systems = make([]*System, 0)
	}
	systems, ok := iniconfig.SectionMap[KEY_SYSTEM_TAG]
	if !ok {
		log.Trace("failed to access iniconfig.SectionMap", KEY_SYSTEM_TAG)
	}
	for _, system := range systems {
		system_to_add := System{}

		system_to_add.Nickname = system.ParamMap[KEY_NICKNAME][0].First.AsString()

		if len(system.ParamMap[KEY_FILE]) > 0 {
			system_to_add.File = PathCreate(system.ParamMap[KEY_FILE][0].First.AsString())
		}

		frelconfig.AddSystem(&system_to_add)
	}

	return frelconfig, iniconfig
}
