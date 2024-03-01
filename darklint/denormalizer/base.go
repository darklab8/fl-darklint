package denormalizer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped/systems_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-darklint/darklint/settings/logus"
	"github.com/darklab8/go-typelog/typelog"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/market_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
)

type DenormalizedBaseGood struct {
	name             string
	recycleCandidate string
}

func NewBaseDenormalizer() *BaseDenormalizer {
	return &BaseDenormalizer{
		baseGoods: make(map[string]*DenormalizedBaseGood),
	}
}

type BaseDenormalizer struct {
	baseGoods map[string]*DenormalizedBaseGood
}

func (d *BaseDenormalizer) Read(parsed *configs_mapped.MappedConfigs) *BaseDenormalizer {
	d.baseGoods = make(map[string]*DenormalizedBaseGood)

	for _, base := range parsed.Universe_config.Bases {
		d.baseGoods[base.Nickname.Get()] = &DenormalizedBaseGood{}
	}

	d.ReadBaseNames(parsed.MarketShips, parsed.Universe_config, parsed.Infocards)
	d.ReadRecycle(parsed.MarketShips, parsed.Universe_config, parsed.Systems)
	return d
}

func (d *BaseDenormalizer) Write(parsed *configs_mapped.MappedConfigs) {
	d.MarketWrite(parsed.MarketCommidities)
	d.MarketWrite(parsed.MarketShips)
	d.MarketWrite(parsed.MarketMisc)
	d.UniverseWrite(parsed.Universe_config)
}

const (
	// denormalized keys
	BASE_KEY_NAME    = "name"
	BASE_KEY_RECYCLE = "is_recycle_candidate"
)

func (denormalizer *BaseDenormalizer) MarketWrite(marketconfig *market_mapped.Config) {

	for _, base_good := range marketconfig.BaseGoods {
		base_denormalized_name := semantic.NewString(base_good.Render(), BASE_KEY_NAME, semantic.TypeComment, inireader.OPTIONAL_p)

		if denorm_base_good, ok := denormalizer.baseGoods[base_good.Base.Get()]; ok {
			base_denormalized_name.Set(denorm_base_good.name)
		}

		base_denorm_RecycleCandidate := semantic.NewString(base_good.Render(), BASE_KEY_RECYCLE, semantic.TypeComment, inireader.OPTIONAL_p)

		if denorm_base_good, ok := denormalizer.baseGoods[base_good.Base.Get()]; ok {
			if denorm_base_good.recycleCandidate != "" {
				base_denorm_RecycleCandidate.Set(denormalizer.baseGoods[base_good.Base.Get()].recycleCandidate)
			} else {
				base_denorm_RecycleCandidate.Delete()
			}
		}
	}
}

func (denormalizer *BaseDenormalizer) ReadBaseNames(marketconfig *market_mapped.Config, universeConfig *universe_mapped.Config, infocards *infocard_mapped.Config) {
	for _, base := range universeConfig.Bases {

		base_strid_name := int(base.StridName.Get())
		record, ok := infocards.RecordsMap[base_strid_name]
		if !ok {
			logus.Log.Error("failed to get record from infocards.txt for base_strid_name", typelog.Int("base_strid_name", base_strid_name), typelog.Any("base_nickname", base.Nickname.Get()), typelog.Any("base_nickname", base.File.Get()))
			continue
		}

		// Fallback reading from dll is not able to recognize kind Yet
		// if (*record).Kind != "NAME" {
		// 	logus.Log.Fatal("incorrect type in infocards, expected NAME for base_strid_name: ", base_strid_name, " record: ", *record, " base: ", base)
		// }

		denormalizer.baseGoods[base.Nickname.Get()].name = (*record).Content
	}
}

var system_for_recycled_bases = [...]string{"ga13", "fp7"}

func (denormalizer *BaseDenormalizer) ReadRecycle(marketconfig *market_mapped.Config, universeConfig *universe_mapped.Config, systems *systems_mapped.Config) {
	for _, base := range universeConfig.Bases {

		system, ok := systems.SystemsMap.MapGetValue(base.System.Get())
		if !ok {
			logus.Log.Fatal("base is leading to non existent system", typelog.String("base_name", base.Nickname.Get()), typelog.String("system", base.System.Get()))
			continue
		}

		var recycle_builder strings.Builder

		_, ok = system.BasesByBase.MapGetValue(base.Nickname.Get())
		if !ok {
			recycle_builder.WriteString(fmt.Sprintf("%s base_good.base=%s not in universe.ini->Base.system->System.file->%s | ", errors.New("base is not present in a system"), base.Nickname.Get(), system.Nickname))
			// logus.Log.Fatal("base ", base_good.Base, " is not present in system ", system.Nickname, " potential crash situation")
		}

		for _, recycle_system := range system_for_recycled_bases {
			if base.System.Get() == recycle_system {
				recycle_builder.WriteString(fmt.Sprintf("universe.ini->Base.system=%s in [%v]", recycle_system, system_for_recycled_bases))
			}
		}

		recycleCandidate := recycle_builder.String()
		denormalizer.baseGoods[base.Nickname.Get()].recycleCandidate = recycleCandidate

	}
}

func (denormalizer *BaseDenormalizer) UniverseWrite(universeConfig *universe_mapped.Config) {

	universeConfig.BasesMap.Foreach(func(base_nickname universe_mapped.BaseNickname, base *universe_mapped.Base) {
		base_name := semantic.NewString(base.Render(), BASE_KEY_NAME, semantic.TypeComment, inireader.OPTIONAL_p)

		if denorm_base_good, ok := denormalizer.baseGoods[string(base_nickname)]; ok {
			if denorm_base_good.name != "" {
				base_name.Set(denorm_base_good.name)
			}

		}

		base_recycle_candidate := semantic.NewString(base.Render(), BASE_KEY_RECYCLE, semantic.TypeComment, inireader.OPTIONAL_p)
		if denorm_base_good, ok := denormalizer.baseGoods[string(base_nickname)]; ok {
			if denorm_base_good.recycleCandidate != "" {
				base_recycle_candidate.Set(denormalizer.baseGoods[string(base_nickname)].recycleCandidate)
			} else {
				base_recycle_candidate.Delete()
			}
		}
	})
}
