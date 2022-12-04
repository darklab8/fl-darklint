/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/settings"
	"darktool/tools/parser/freelancer/data/equipment/market"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"
	"darktool/tools/parser/parserutils/filefind"

	log "github.com/sirupsen/logrus"
)

func Parse(file1path string) {
	log.Info("Parse START for FreelancerFolderLocation=", file1path)
	filesystem := filefind.FindConfigs(file1path)

	universe_config := universe.Config{}
	_, iniconfig := universe_config.Read(filesystem.GetFile(universe.FILENAME))
	iniconfig.Write(filesystem.GetFile(universe.FILENAME)).WriteLines()

	info_config := service.Config{}
	info_config.Read(filesystem.GetFile(service.FILENAME, service.FILENAME_FALLBACK))

	market_ships_config := market.Config{}
	market_ships_config.Read(filesystem.GetFile(market.FILENAME_SHIPS))
	market_ships_config.UpdateWithBasenames(&universe_config, &info_config)
	market_ships_config.Write(filesystem.GetFile(market.FILENAME_SHIPS)).WriteLines()

	market_commodities := market.Config{}
	market_commodities.Read(filesystem.GetFile(market.FILENAME_COMMODITIES))
	market_commodities.UpdateWithBasenames(&universe_config, &info_config)
	market_commodities.Write(filesystem.GetFile(market.FILENAME_COMMODITIES)).WriteLines()

	// TODO implement preserving comments :|
	// market_misc := market1ships.Config{}
	// market_misc.Read(filesystem.GetFile(market1ships.FILENAME_MISC))
	// market_misc.UpdateWithBasenames(&universe_config, &info_config)
	// market_misc.Write(filesystem.GetFile(market1ships.FILENAME_MISC)).WriteLines()

	log.Info("Parse OK for FreelancerFolderLocation=", file1path)
}

func Run() {
	Parse(settings.FreelancerFolderLocation)
}
