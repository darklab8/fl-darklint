package universe_format

import (
	"strings"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
)

type UniverseFormatter struct {
	universe_config *universe_mapped.Config
}

func NewUniverseFormatter(universe_config *universe_mapped.Config) *UniverseFormatter {
	return &UniverseFormatter{universe_config: universe_config}
}

func (f *UniverseFormatter) Format() {
	for _, base := range f.universe_config.Bases {
		base.Nickname.Set(strings.ToLower(base.Nickname.Get()))
		base.System.Set(strings.ToLower(base.System.Get()))
		base.File.Set(strings.ToLower(base.File.Get()))
	}

	for _, system := range f.universe_config.Systems {
		system.Nickname.Set(strings.ToLower(system.Nickname.Get()))
		system.Msg_id_prefix.Set(strings.ToLower(system.Msg_id_prefix.Get()))

		if system.File.Get() != "" {
			system.File.Set(strings.ToLower(system.File.Get()))
		}
	}
}
