package market1ships

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"

	log "github.com/sirupsen/logrus"
)

type MarketGood struct {
	Name   string
	Values []inireader.ValueNumber
}

type BaseGood struct {
	// Name  string // implement commented out variables
	Base  string
	Name  string
	Goods []*MarketGood
}

type Config struct {
	Intro     []string
	BaseGoods []*BaseGood
}

var LoadedConfig Config

const BaseGoodType = "[BaseGood]"

func Read(input_file utils.File) Config {
	var frelconfig Config

	iniconfig := inireader.INIFileRead(input_file)

	for _, section := range iniconfig.Sections {
		if section.Type != BaseGoodType {
			log.Fatalf("%v != %v", section.Type, BaseGoodType)
		}
		if len(section.Params) == 0 {
			continue
		}
		current_base_good := BaseGood{}
		frelconfig.BaseGoods = append(frelconfig.BaseGoods, &current_base_good)

		if !utils.IsLower(section.ParamMap["base"][0].First.AsString()) {
			log.Warn("market_ships, base: ", section.ParamMap["base"][0].First, " not in a lower case, autofixing")
		}
		current_base_good.Base = string(section.ParamMap["base"][0].First.(inireader.ValueString).ToLowerValue())

		good_params, ok := section.ParamMap["MarketGood"]
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

	return frelconfig
}

func Load() {
	file := utils.File{Filepath: filefind.FreelancerFolder.Hashmap["market_ships.ini"].Filepath}
	LoadedConfig = Read(file)
	log.Info("OK market_ships.ini is parsed to specialized data structs")
}
