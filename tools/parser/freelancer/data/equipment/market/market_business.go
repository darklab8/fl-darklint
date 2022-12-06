/*
Business rule to market_ships.txt
*/
package market

import (
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/data/universe/systems"
	"darktool/tools/parser/freelancer/infocard"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (frelconfig *Config) UpdateWithBasenames(universeConfig *universe.Config, infocards *infocard.Config) {
	for _, base_good := range frelconfig.BaseGoods {
		key := universe.BaseNickname(base_good.Base.Get())
		base, ok := universeConfig.BasesMap[key]
		if !ok {
			log.Warn("failed to get key: ", key, " from universe.txt for base during updating market_ships.txt, attempting to fix")
		}

		base_strid_name := int(base.StridName.Get())
		record, ok := infocards.RecordsMap[base_strid_name]
		if !ok {
			log.Fatal("failed to get record from infocardts.txt, with id: ", base_strid_name, " for base: ", base, " record = ", record)
		}

		if (*record).Type() != "NAME" {
			log.Fatal("incorrect type in infocards, expected NAME for base_strid_name: ", base_strid_name, " record: ", *record, " base: ", base)
		}

		base_good.Name.Set((*record).Content())
	}
}

var system_for_recycled_bases = [...]string{"ga13", "fp7"}

func (frelconfig *Config) UpdateWithRecycle(universeConfig *universe.Config, systems *systems.Config) {
	for _, base_good := range frelconfig.BaseGoods {
		universe_base, ok := universeConfig.BasesMap[universe.BaseNickname(base_good.Base.Get())]
		if !ok {
			log.Fatal("base_good=", base_good.Base, "is not having universe base data")
		}

		system, ok := systems.SystemsMap[universe_base.System.Get()]
		if !ok {
			log.Fatal("base ", universe_base.Nickname, "is leading to non existent system", universe_base.System)
		}

		var recycle_builder strings.Builder

		universe_system, ok := universeConfig.SystemMap[universe.SystemNickname(system.Nickname)]
		_, ok = system.BasesByBase[base_good.Base.Get()]
		if !ok {
			recycle_builder.WriteString(fmt.Sprintf("base_good.base=%v not in universe.ini->Base.system->System.file->%v | ", base_good.Base, universe_system.File))
			// log.Fatal("base ", base_good.Base, " is not present in system ", system.Nickname, " potential crash situation")
		}

		for _, recycle_system := range system_for_recycled_bases {
			if universe_base.System.Get() == recycle_system {
				recycle_builder.WriteString(fmt.Sprintf("universe.ini->Base.system=%s in [%v]", recycle_system, system_for_recycled_bases))
			}
		}

		recycleCandidate := recycle_builder.String()
		if recycleCandidate != "" {
			base_good.RecycleCandidate.Set(recycleCandidate)
		} else {
			base_good.RecycleCandidate.Delete()
		}

	}
}
