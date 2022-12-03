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

// type MarketGood struct {
// 	Name   string
// 	Values []inireader.ValueNumber
// }

// type BaseGood struct {
// 	// Name  string // implement commented out variables
// 	Base  string
// 	Goods []*MarketGood
// }

// type MarketShips struct {
// 	Intro      []string
// 	Base_goods []*BaseGood
// }

type Config struct {
}

// var LoadedMarketShips MarketShips

// const BaseGoodType = "[BaseGood]"

const filename = "universe.ini"

func Read(input_file utils.File) Config {
	var frelconfig Config

	_ = inireader.INIFileRead(input_file)

	// for _, section := range iniconfig.Sections {
	// 	if section.Type != BaseGoodType {
	// 		log.Fatalf("%v != %v", section.Type, BaseGoodType)
	// 	}
	// 	if len(section.Params) == 0 {
	// 		continue
	// 	}
	// 	current_base_good := BaseGood{}
	// 	frelconfig.Base_goods = append(frelconfig.Base_goods, &current_base_good)

	// 	current_base_good.Base = string(section.ParamMap["base"][0].First.(inireader.ValueString))

	// 	good_params, ok := section.ParamMap["MarketGood"]
	// 	if ok {
	// 		for _, good_param := range good_params {

	// 			good := MarketGood{}
	// 			good.Name = string(good_param.First.(inireader.ValueString))

	// 			for _, value := range good_param.Values[1:] {
	// 				good.Values = append(good.Values, value.(inireader.ValueNumber))
	// 			}
	// 			current_base_good.Goods = append(current_base_good.Goods, &good)

	// 		}
	// 	}
	// }

	return frelconfig
}

func Load() {
	file := utils.File{Filepath: filefind.FreelancerFolder.Hashmap[filename].Filepath}
	Read(file)
	log.Info("OK universe.ini is parsed to specialized data structs")
}
