/*
Business rule to market_ships.txt
*/
package market

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"

	log "github.com/sirupsen/logrus"
)

func (frelconfig *Config) UpdateWithBasenames(universeConfig *universe.Config, infocards *service.Config) {
	for _, base_good := range frelconfig.BaseGoods {
		key := universe.BaseNickname(base_good.Base)
		base, ok := universeConfig.BasesMap[key]
		if !ok {
			log.Warn("failed to get key: ", key, " from universe.txt for base during updating market_ships.txt, attempting to fix")
		}

		base_strid_name := int(base.StridName)
		record, ok := infocards.RecordsMap[base_strid_name]
		if !ok {
			log.Fatal("failed to get record from infocardts.txt, with id: ", base_strid_name, " for base: ", base, " record = ", record)
		}

		if (*record).Type() != "NAME" {
			log.Fatal("incorrect type in infocards, expected NAME for base_strid_name: ", base_strid_name, " record: ", *record, " base: ", base)
		}

		base_good.Name = (*record).Content()
	}
}
