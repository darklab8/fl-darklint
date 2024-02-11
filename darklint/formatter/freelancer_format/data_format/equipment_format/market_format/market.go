package market_format

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/market_mapped"
)

type ConfigFormatter struct {
	config *market_mapped.Config
}

func NewFormatter(config *market_mapped.Config) *ConfigFormatter {
	return &ConfigFormatter{config: config}
}

func (f *ConfigFormatter) Format() {
	for _, base_good := range f.config.BaseGoods {
		base_good.Base.Set(strings.ToLower(base_good.Base.Get()))
	}
}
