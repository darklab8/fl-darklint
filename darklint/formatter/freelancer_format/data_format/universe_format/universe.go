package universe_format

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
)

type ConfigFormatter struct {
	config *universe_mapped.Config
}

func NewFormatter(universe_config *universe_mapped.Config) *ConfigFormatter {
	return &ConfigFormatter{config: universe_config}
}

func (f *ConfigFormatter) Format() {
	for _, base := range f.config.Bases {
		base.Nickname.Set(strings.ToLower(base.Nickname.Get()))
		base.System.Set(strings.ToLower(base.System.Get()))
		base.File.Set(strings.ToLower(base.File.Get()))
	}

	for _, system := range f.config.Systems {
		system.Nickname.Set(strings.ToLower(system.Nickname.Get()))
		system.Msg_id_prefix.Set(strings.ToLower(system.Msg_id_prefix.Get()))

		if system.File.Get() != "" {
			system.File.Set(strings.ToLower(system.File.Get()))
		}
	}
}
