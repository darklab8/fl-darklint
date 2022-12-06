package market

import (
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"strings"
)

// ORM values

const (
	TypeComment = true
	TypeVisible = false
)

type SemanticValue struct {
	section   *inireader.Section
	key       string
	optional  bool
	isComment bool
}

type SemanticString struct {
	SemanticValue
}

func (s *SemanticString) Map(section *inireader.Section, key string, isComment bool, optional bool) *SemanticString {
	s.section = section
	s.key = key
	s.optional = optional
	s.isComment = isComment
	return s
}

func (s *SemanticString) Get() string {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return ""
	}
	return s.section.ParamMap[s.key][0].First.AsString()
}

func (s *SemanticString) Set(value string) {
	processed_value := inireader.UniParseStr(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParam(s.key, (&inireader.Param{IsComment: s.isComment}).AddValue(processed_value))
	}
	// implement SetValue in Section
	s.section.ParamMap[s.key][0].First = processed_value
	s.section.ParamMap[s.key][0].Values[0] = processed_value
}

func (s *SemanticString) Delete() {
	delete(s.section.ParamMap, s.key)
	for index, param := range s.section.Params {
		if param.Key == s.key {
			s.section.Params = append(s.section.Params[:index], s.section.Params[index+1:]...)
			break
		}
	}
}

// ORM Model

type SemanticModel struct {
	section *inireader.Section
}

func (s *SemanticModel) Map(section *inireader.Section) {
	s.section = section
}

func (s *SemanticModel) Render() *inireader.Section {
	return s.section
}

// Market

// Not implemented. Create SemanticMultiKeyValue
type MarketGood struct {
	SemanticModel
	Name *SemanticString
	// Values SemanticIntArray
}

type BaseGood struct {
	SemanticModel
	Base *SemanticString
	// TODO Goods          *SemanticMultiKey[MarketGood] (GetAll)
	Name             *SemanticString // denormalized always disabled param
	RecycleCandidate *SemanticString // denormalized always disabled param
}

type Config struct {
	BaseGoods []*BaseGood
	Comments  []string
}

const (
	FILENAME_SHIPS            = "market_ships.ini"
	FILENAME_COMMODITIES      = "market_commodities.ini"
	FILENAME_MISC             = "market_misc.ini"
	BaseGoodType              = "[BaseGood]"
	KEY_RECYCLE               = "is_recycle_candidate"
	KEY_MISSMATCH_SYSTEM_FILE = "missmatched_universe_system_and_file"
	KEY_MARKET_GOOD           = "marketgood"
	KEY_BASE                  = "base"
	KEY_NAME                  = "name"
)

func (frelconfig *Config) Read(input_file *utils.File) *Config {
	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.BaseGoods = make([]*BaseGood, 0)

	for _, section := range iniconfig.Sections {
		base_to_add := &BaseGood{}
		base_to_add.Map(section)
		base_to_add.Base = (&SemanticString{}).Map(section, KEY_BASE, TypeVisible, inireader.REQUIRED_p)
		base_to_add.Name = (&SemanticString{}).Map(section, KEY_NAME, TypeComment, inireader.OPTIONAL_p)
		base_to_add.RecycleCandidate = (&SemanticString{}).Map(section, KEY_BASE, TypeVisible, inireader.OPTIONAL_p)
		frelconfig.BaseGoods = append(frelconfig.BaseGoods, base_to_add)

		base_to_add.Base.Set(strings.ToLower(base_to_add.Base.Get()))
	}
	frelconfig.Comments = iniconfig.Comments
	return frelconfig
}

func (frelconfig *Config) Write(output_file *utils.File) *utils.File {

	inifile := inireader.INIFile{}
	inifile.File = output_file
	inifile.Comments = frelconfig.Comments

	for _, baseGood := range frelconfig.BaseGoods {
		inifile.Sections = append(inifile.Sections, baseGood.Render())
	}

	inifile.Write(output_file)
	return inifile.File
}
