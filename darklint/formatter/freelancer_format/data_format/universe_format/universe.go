package universe_format

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
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
		base.File.Set(utils_types.FilePath(strings.ToLower(base.File.Get().ToString())))
	}

	for _, system := range f.config.Systems {
		system.Nickname.Set(strings.ToLower(system.Nickname.Get()))
		system.Msg_id_prefix.Set(strings.ToLower(system.Msg_id_prefix.Get()))

		if system.File.Get() != "" {
			system.File.Set(utils_types.FilePath(strings.ToLower(system.File.Get().ToString())))
		}
	}
}
