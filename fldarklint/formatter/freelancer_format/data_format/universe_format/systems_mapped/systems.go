package systems_mapped

import (
	"strings"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped/systems_mapped"
)

type ConfigFormatter struct {
	config *systems_mapped.Config
}

func NewFormatter(config *systems_mapped.Config) *ConfigFormatter {
	return &ConfigFormatter{config: config}
}

func (f *ConfigFormatter) Format() {
	for _, system := range f.config.Systems {
		for _, base := range system.Bases {
			base.Nickname.Set(strings.ToLower(base.Nickname.Get()))
			base.Base.Set(strings.ToLower(base.Base.Get()))
			if base.DockWith.Get() != "" {
				base.DockWith.Set(strings.ToLower(base.DockWith.Get()))
			}
		}
	}
}
