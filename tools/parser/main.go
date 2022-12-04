/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/tools/parser/freelancer/data/equipment/market1ships"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/parser/parserutils/inireader"

	log "github.com/sirupsen/logrus"
)

func Run() {
	filefind.Load()
	market1ships.Load()
	universe.Load()
	service.Load()

	for _, base_good := range market1ships.LoadedConfig.BaseGoods {
		key := universe.BaseNickname(base_good.Base)
		base, ok := universe.Loaded.BasesMap[key]
		if !ok {
			log.Warn("failed to get key: ", key, " from universe.txt for base during updating market_ships.txt, attempting to fix")
		}

		base_strid_name := int(base.StridName.(inireader.ValueNumber).Value)
		record, ok := service.LoadedInfocards.RecordsMap[base_strid_name]
		if !ok {
			log.Fatal("failed to get record from infocardts.txt, with id: ", base_strid_name, " for base: ", base, " record = ", record)
		}

		if (*record).Type() != "NAME" {
			log.Fatal("incorrect type in infocards, expected NAME for base_strid_name: ", base_strid_name, " record: ", *record, " base: ", base)
		}

		base_good.Name = (*record).Content()
	}
}
