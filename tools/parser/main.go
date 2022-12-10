/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/tools/parser/freelancer/data/equipment/market"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/data/universe/systems"
	"darktool/tools/parser/freelancer/infocard"
	"darktool/tools/parser/parserutils/filefind"

	log "github.com/sirupsen/logrus"
)

type Parsed struct {
	Universe_config     *universe.Config
	Info_config         *infocard.Config
	Systems             *systems.Config
	Market_ships_config *market.Config
	Market_commodities  *market.Config
	Market_misc         *market.Config
}

func (p *Parsed) Read(file1path string) *Parsed {
	log.Info("Parse START for FreelancerFolderLocation=", file1path)
	filesystem := filefind.FindConfigs(file1path)

	p.Universe_config = (&universe.Config{}).Read(filesystem.GetFile(universe.FILENAME))
	p.Info_config = (&infocard.Config{}).Read(filesystem.GetFile(infocard.FILENAME, infocard.FILENAME_FALLBACK))
	p.Systems = (&systems.Config{}).Read(p.Universe_config, filesystem)
	p.Market_ships_config = (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_SHIPS))
	p.Market_commodities = (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_COMMODITIES))
	p.Market_misc = (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_MISC))

	log.Info("Parse OK for FreelancerFolderLocation=", file1path)

	return p
}

func (p *Parsed) Write(dry_run bool) {
	// write
	p.Universe_config.Write().WriteLines(dry_run)
	p.Market_ships_config.Write().WriteLines(dry_run)
	p.Market_commodities.Write().WriteLines(dry_run)
	p.Market_misc.Write().WriteLines(dry_run)
	system_files := p.Systems.Write()
	for _, system_file := range system_files {
		system_file.WriteLines(dry_run)
	}

}
