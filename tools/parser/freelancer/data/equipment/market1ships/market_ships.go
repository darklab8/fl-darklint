package market1ships

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
)

type MarketGood struct {
	Name   string
	Values []float32
}

type BaseGood struct {
	Name  string
	Base  string
	Goods []MarketGood
}

type MarketShips struct {
	Intro      []string
	Base_goods []*BaseGood
}

var LoadedMarketShips MarketShips

func Read(input_file utils.File) MarketShips {
	var frelconfig MarketShips

	iniconfig := inireader.INIFileRead(input_file)
	_ = iniconfig

	return frelconfig
}

func Load() {
	file := utils.File{Filepath: filefind.FreelancerFolder.Hashmap["market_ships.ini"].Filepath}
	LoadedMarketShips = Read(file)
}
