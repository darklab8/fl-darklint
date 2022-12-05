package market

import (
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"

	log "github.com/sirupsen/logrus"
)

type MarketGood struct {
	Name   string
	Values []inireader.ValueNumber
}

type BaseGood struct {
	Base             string
	Name             string // denormalized always disabled param
	Goods            []*MarketGood
	RecycleCandidate string // denormalized always disabled param
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

	if frelconfig.BaseGoods == nil {
		frelconfig.BaseGoods = make([]*BaseGood, 0)
	}

	for _, section := range iniconfig.Sections {
		if section.Type != BaseGoodType {
			log.Fatalf("%v != %v", section.Type, BaseGoodType)
		}
		if len(section.Params) == 0 {
			continue
		}
		current_base_good := BaseGood{}
		if current_base_good.Goods == nil {
			current_base_good.Goods = make([]*MarketGood, 0)
		}
		frelconfig.BaseGoods = append(frelconfig.BaseGoods, &current_base_good)

		current_base_good.Name = section.GetParamStrToLower(KEY_NAME, inireader.OPTIONAL_p)
		current_base_good.Base = section.GetParamStrToLower(KEY_BASE, inireader.REQUIRED_p)

		good_params, ok := section.ParamMap[KEY_MARKET_GOOD]
		if ok {
			for _, good_param := range good_params {

				good := MarketGood{}
				good.Name = string(good_param.First.(inireader.ValueString))

				for _, value := range good_param.Values[1:] {
					good.Values = append(good.Values, value.(inireader.ValueNumber))
				}
				current_base_good.Goods = append(current_base_good.Goods, &good)

			}
		}
	}
	frelconfig.Comments = iniconfig.Comments
	return frelconfig
}

func (frelconfig *Config) Write(output_file *utils.File) *utils.File {

	inifile := inireader.INIFile{}
	inifile.File = output_file
	inifile.Comments = frelconfig.Comments

	for _, baseGood := range frelconfig.BaseGoods {
		section := inireader.Section{}
		section.Type = BaseGoodType
		section.AddParam(KEY_BASE, (&inireader.Param{}).AddValue(inireader.UniParseStr(baseGood.Base)))

		// Optional
		if baseGood.Name != "" {
			section.AddParam(KEY_NAME, (&inireader.Param{IsComment: true}).AddValue(inireader.UniParseStr(baseGood.Name)))
		}
		if baseGood.RecycleCandidate != "" {
			section.AddParam(KEY_RECYCLE, (&inireader.Param{IsComment: true}).AddValue(inireader.UniParseStr(baseGood.RecycleCandidate)))
		}

		for _, param := range baseGood.Goods {
			market_good := inireader.Param{Key: KEY_MARKET_GOOD, IsComment: false}

			market_good.AddValue(inireader.ValueString(param.Name))
			for _, value := range param.Values {
				market_good.AddValue(value)
			}
			section.Params = append(section.Params, &market_good)
		}
		inifile.Sections = append(inifile.Sections, &section)
	}

	inifile.Write(output_file)
	return inifile.File
}
