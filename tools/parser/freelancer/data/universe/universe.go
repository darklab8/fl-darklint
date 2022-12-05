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

// terrain_tiny = nonmineable_asteroid90
// terrain_sml = nonmineable_asteroid60
// terrain_mdm = nonmineable_asteroid90
// terrain_lrg = nonmineable_asteroid60
// terrain_dyna_01 = mineable1_asteroid10
// terrain_dyna_02 = mineable1_asteroid10

var KEY_BASE_TERRAINS = [...]string{"terrain_tiny", "terrain_sml", "terrain_mdm", "terrain_lrg", "terrain_dyna_01", "terrain_dyna_02"}

const (
	FILENAME      = "universe.ini"
	KEY_BASE_TAG  = "[Base]"
	KEY_NICKNAME  = "nickname"
	KEY_STRIDNAME = "strid_name"
	KEY_SYSTEM    = "system"
	KEY_FILE      = "file"

	KEY_BASE_BGCS = "BGCS_base_run_by"

	KEY_SYSTEM_TAG           = "[system]"
	KEY_SYSTEM_MSG_ID_PREFIX = "msg_id_prefix"
	KEY_SYSTEM_VISIT         = "visit"
	KEY_SYSTEM_IDS_INFO      = "ids_info"
	KEY_SYSTEM_NAVMAPSCALE   = "NavMapScale"
	KEY_SYSTEM_POS           = "pos"

	KEY_TIME_TAG     = "[Time]"
	KEY_TIME_SECONDS = "seconds_per_day"
)

type Time struct {
	Seconds_per_day int
}

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

	Terrains map[string]string
}

type BaseNickname string

type SystemNickname string

type System struct {
	Nickname      string
	Pos           [2]inireader.ValueNumber
	Msg_id_prefix string
	Visit         int
	Strid_name    int
	Ids_info      int
	File          Path
	NavMapScale   inireader.ValueNumber
}

type Config struct {
	Bases    []*Base
	BasesMap map[BaseNickname]*Base //key is

	Systems   []*System
	SystemMap map[SystemNickname]*System //key is

	TimeSeconds Time
}

func (frelconfig *Config) AddBase(base_to_add *Base) {
	frelconfig.Bases = append(frelconfig.Bases, base_to_add)
	frelconfig.BasesMap[BaseNickname(base_to_add.Nickname)] = base_to_add
}

func (frelconfig *Config) AddSystem(system_to_add *System) {
	frelconfig.Systems = append(frelconfig.Systems, system_to_add)
	frelconfig.SystemMap[SystemNickname(system_to_add.Nickname)] = system_to_add
}

func (frelconfig *Config) Read(input_file *utils.File) *Config {
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

		if len(base.ParamMap[KEY_BASE_BGCS]) > 0 {
			base_to_add.BGCS_base_run_by = base.ParamMap[KEY_BASE_BGCS][0].First.AsString()
		}

		if base_to_add.Terrains == nil {
			base_to_add.Terrains = make(map[string]string)
		}
		for _, terrain_key := range KEY_BASE_TERRAINS {
			terrain_param, ok := base.ParamMap[terrain_key]
			if ok {
				base_to_add.Terrains[terrain_key] = terrain_param[0].First.AsString()
			}
		}

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

		system_to_add.Nickname = strings.ToLower(system.ParamMap[KEY_NICKNAME][0].First.AsString())

		key_param, ok := system.ParamMap[KEY_SYSTEM_POS]
		if ok {
			if len(key_param[0].Values) < 2 {
				log.Fatal("unable to parse Pos for system, incorrect len", system.Params, system_to_add, "pos=", system.ParamMap[KEY_SYSTEM_POS][0].First.AsString(), "nickname=", system_to_add.Nickname)
			}
			for index, number := range key_param[0].Values {
				system_to_add.Pos[index] = number.(inireader.ValueNumber)
			}
		}

		if len(system.ParamMap[KEY_FILE]) > 0 {
			system_to_add.File = PathCreate(system.ParamMap[KEY_FILE][0].First.AsString())
		}

		if len(system.ParamMap[KEY_SYSTEM_MSG_ID_PREFIX]) > 0 {
			system_to_add.Msg_id_prefix = strings.ToLower(system.ParamMap[KEY_SYSTEM_MSG_ID_PREFIX][0].First.AsString())
		}

		if len(system.ParamMap[KEY_SYSTEM_VISIT]) > 0 {
			visits, err := strconv.Atoi(strings.ToLower(system.ParamMap[KEY_SYSTEM_VISIT][0].First.AsString()))
			if err == nil {
				system_to_add.Visit = visits
			}
		}

		if len(system.ParamMap[KEY_STRIDNAME]) > 0 {
			strid_name, err := strconv.Atoi(strings.ToLower(system.ParamMap[KEY_STRIDNAME][0].First.AsString()))
			if err == nil {
				system_to_add.Strid_name = strid_name
			}
		}

		if len(system.ParamMap[KEY_SYSTEM_IDS_INFO]) > 0 {
			ids_info, err := strconv.Atoi(strings.ToLower(system.ParamMap[KEY_SYSTEM_IDS_INFO][0].First.AsString()))
			if err == nil {
				system_to_add.Ids_info = ids_info
			}
		}

		if len(system.ParamMap[KEY_SYSTEM_NAVMAPSCALE]) > 0 {
			system_to_add.NavMapScale = system.ParamMap[KEY_SYSTEM_NAVMAPSCALE][0].First.(inireader.ValueNumber)
		}

		frelconfig.AddSystem(&system_to_add)
	}

	// TIME
	seconds_per_day, err := strconv.Atoi(iniconfig.SectionMap[KEY_TIME_TAG][0].ParamMap[KEY_TIME_SECONDS][0].First.AsString())
	if err != nil {
		log.Fatal("unable to parse time in universe.ini")
	}
	frelconfig.TimeSeconds = Time{Seconds_per_day: seconds_per_day}

	return frelconfig
}

func (frelconfig *Config) Write(output_file *utils.File) *utils.File {
	inifile := inireader.INIFile{}
	inifile.File = output_file

	time_section := inireader.Section{Type: KEY_TIME_TAG}
	time_section.AddParam(KEY_TIME_SECONDS, (&inireader.Param{Key: KEY_TIME_SECONDS}).AddValue(inireader.UniParseInt(frelconfig.TimeSeconds.Seconds_per_day)))
	inifile.AddSection(KEY_SYSTEM_TAG, &time_section)

	for _, base := range frelconfig.Bases {
		base_to_add := inireader.Section{Type: KEY_BASE_TAG}

		base_to_add.AddParam(KEY_NICKNAME, (&inireader.Param{}).AddValue(inireader.UniParseStr(base.Nickname)))
		base_to_add.AddParam(KEY_SYSTEM, (&inireader.Param{}).AddValue(inireader.UniParseStr(base.System)))
		base_to_add.AddParam(KEY_STRIDNAME, (&inireader.Param{}).AddValue(inireader.UniParseInt(base.StridName)))
		base_to_add.AddParam(KEY_FILE, (&inireader.Param{}).AddValue(inireader.UniParseStr(base.File.WindowsPath())))
		base_to_add.AddParam(KEY_BASE_BGCS, (&inireader.Param{}).AddValue(inireader.UniParseStr(base.BGCS_base_run_by)))

		for terrain_key, terrain_value := range base.Terrains {
			base_to_add.AddParam(terrain_key, (&inireader.Param{}).AddValue(inireader.UniParseStr(terrain_value)))
		}

		inifile.AddSection(KEY_BASE_TAG, &base_to_add)
	}

	for _, system := range frelconfig.Systems {
		system_to_add := inireader.Section{Type: KEY_SYSTEM_TAG}

		system_to_add.AddParam(KEY_NICKNAME, (&inireader.Param{}).AddValue(inireader.UniParseStr(system.Nickname)))
		if system.File.WindowsPath() != "" {
			system_to_add.AddParam(KEY_FILE, (&inireader.Param{}).AddValue(inireader.UniParseStr(system.File.WindowsPath())))
		}
		system_to_add.AddParam(KEY_SYSTEM_POS, (&inireader.Param{}).AddValue(system.Pos[0]).AddValue(system.Pos[1]))
		system_to_add.AddParam(KEY_SYSTEM_MSG_ID_PREFIX, (&inireader.Param{}).AddValue(inireader.UniParseStr(system.Msg_id_prefix)))
		system_to_add.AddParam(KEY_SYSTEM_VISIT, (&inireader.Param{}).AddValue(inireader.UniParseInt(system.Visit)))
		system_to_add.AddParam(KEY_STRIDNAME, (&inireader.Param{}).AddValue(inireader.UniParseInt(system.Strid_name)))
		system_to_add.AddParam(KEY_SYSTEM_IDS_INFO, (&inireader.Param{}).AddValue(inireader.UniParseInt(system.Ids_info)))

		if system.NavMapScale.Value != 0 {

			system_to_add.AddParam(KEY_SYSTEM_NAVMAPSCALE, (&inireader.Param{}).AddValue(system.NavMapScale))
		}

		inifile.AddSection(KEY_SYSTEM_TAG, &system_to_add)

	}

	log.Info("formed inifile for universe.ini for writing")
	inifile.Write(output_file)
	log.Info("wrote inifile for universe.ini back to disk")
	return inifile.File
}
