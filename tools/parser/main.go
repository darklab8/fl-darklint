/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/settings"
	"darktool/tools/parser/freelancer/data/equipment/market1ships"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"
	"darktool/tools/parser/parserutils/filefind"

	log "github.com/sirupsen/logrus"
)

func Parse(file1path string) {
	log.Info("Parse START for FreelancerFolderLocation=", file1path)
	filesystem := filefind.FindConfigs(file1path)

	output_config := filesystem.GetFile(market1ships.Filename)

	market_config := market1ships.Config{}
	market_config.Read(filesystem.GetFile(market1ships.Filename))

	universe_config := universe.Config{}
	universe_config.Read(filesystem.GetFile(universe.FILENAME))

	info_config := service.Config{}
	info_config.Read(filesystem.GetFile(service.FILENAME, service.FILENAME_FALLBACK))

	market_config.UpdateWithBasenames(&universe_config, &info_config)
	market_config.Write(output_config)

	output_config.WriteLines()
	log.Info("Parse OK for FreelancerFolderLocation=", file1path)
}

func Run() {
	Parse(settings.FreelancerFolderLocation)
}
