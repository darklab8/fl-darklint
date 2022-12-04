package market1ships

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type MarketGood struct {
	Name   string
	Values []inireader.ValueNumber
}

type BaseGood struct {
	Base               string
	Name               string
	Goods              []*MarketGood
	isRecycleCandidate bool
}

type Config struct {
	BaseGoods []*BaseGood
	Comments  []string
}

var LoadedConfig *Config

const (
	Filename     = "market_ships.ini"
	BaseGoodType = "[BaseGood]"
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

		// Add Name and Recycle Candidate
		name, ok := section.ParamMap["name"]
		if ok {
			if len(name) > 0 {
				current_base_good.Name = name[0].First.AsString()
			}
		}

		// Add isRecycleCandidate
		isRecycleCandidate, ok := section.ParamMap["isRecycleCandidate"]
		if ok {
			value := isRecycleCandidate[0].First.AsString()
			bool_value, _ := strconv.ParseBool(value)
			current_base_good.isRecycleCandidate = bool_value
		}

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
	frelconfig.Comments = iniconfig.Comments
	return frelconfig
}

func Load() {
	file := &utils.File{Filepath: filefind.FreelancerFolder.Hashmap[Filename].Filepath}
	config := Config{}
	LoadedConfig = config.Read(file)
	log.Info(fmt.Sprintf("OK file.Filepath=%v market_ships.ini is parsed to specialized data structs", file.Filepath))
}

func Unload() *utils.File {
	file := &utils.File{Filepath: filefind.FreelancerFolder.Hashmap[Filename].Filepath}
	LoadedConfig.Write(file)
	log.Info(fmt.Sprintf("OK file.Filepath=%v market_ships.ini is written back", file.Filepath))
	return file
}

func (frelconfig Config) Write(output_file *utils.File) *utils.File {

	inifile := inireader.INIFile{}
	inifile.File = output_file
	inifile.Comments = frelconfig.Comments

	for _, baseGood := range frelconfig.BaseGoods {
		section := inireader.Section{}
		section.Type = BaseGoodType

		base_param := inireader.Param{Key: "base", IsComment: false}
		base_param.AddValue(inireader.ValueString(baseGood.Base))
		section.Params = append(section.Params, &base_param)

		if baseGood.Name != "" {
			name := inireader.Param{Key: "name", IsComment: true}
			name.AddValue(inireader.ValueString(baseGood.Name))
			section.Params = append(section.Params, &name)
		}

		if baseGood.isRecycleCandidate {
			recycle := inireader.Param{Key: "isRecycleCandidate", IsComment: true}
			recycle.AddValue(inireader.ValueBool(baseGood.isRecycleCandidate))
			section.Params = append(section.Params, &recycle)
		}

		for _, param := range baseGood.Goods {
			market_good := inireader.Param{Key: "MarketGood", IsComment: false}

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
