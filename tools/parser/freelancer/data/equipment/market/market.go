package market

import (
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/parser/parserutils/semantic"
	"darktool/tools/utils"
	"strings"
)

// Not implemented. Create SemanticMultiKeyValue
type MarketGood struct {
	semantic.Model
	Name *semantic.String
	// Values SemanticIntArray
}

type BaseGood struct {
	semantic.Model
	Base *semantic.String
	// TODO Goods          *SemanticMultiKey[MarketGood] (GetAll)
	Name             *semantic.String // denormalized always disabled param
	RecycleCandidate *semantic.String // denormalized always disabled param
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
		base_to_add.Base = (&semantic.String{}).Map(section, KEY_BASE, semantic.TypeVisible, inireader.REQUIRED_p)
		base_to_add.Name = (&semantic.String{}).Map(section, KEY_NAME, semantic.TypeComment, inireader.OPTIONAL_p)
		base_to_add.RecycleCandidate = (&semantic.String{}).Map(section, KEY_RECYCLE, semantic.TypeComment, inireader.OPTIONAL_p)
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
