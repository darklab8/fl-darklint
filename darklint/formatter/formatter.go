package formatter

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/darklab8/fl-darklint/darklint/formatter/freelancer_format/data_format/equipment_format/market_format"
	"github.com/darklab8/fl-darklint/darklint/formatter/freelancer_format/data_format/universe_format"
	"github.com/darklab8/fl-darklint/darklint/formatter/freelancer_format/data_format/universe_format/systems_mapped"
	"github.com/darklab8/fl-darklint/darklint/settings"
	"github.com/darklab8/fl-darklint/darklint/settings/logus"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type ConfigFormatter interface {
	Format()
}

type Formatter struct {
	configs    *configs_mapped.MappedConfigs
	formatters []ConfigFormatter
}

func NewFormatter(configs *configs_mapped.MappedConfigs) *Formatter {
	f := &Formatter{
		configs: configs,
	}
	f.formatters = append(f.formatters, universe_format.NewFormatter(configs.Universe_config))
	f.formatters = append(f.formatters, systems_mapped.NewFormatter(configs.Systems))

	f.formatters = append(f.formatters,
		market_format.NewFormatter(f.configs.Market_commodities),
		market_format.NewFormatter(f.configs.Market_commodities),
		market_format.NewFormatter(f.configs.Market_misc),
	)
	return f
}

func (f *Formatter) Format() {
	for _, formatter := range f.formatters {
		formatter.Format()
	}
}

func Run(is_dry_run configs_mapped.IsDruRun) {

	_, err := os.Stat(filepath.Join(settings.FreelancerFreelancerLocation, "DATA"))

	fmt.Println(err)
	if os.IsNotExist(err) {
		logus.Log.Fatal("freelancer folder is not detected because DATA folder was not found", utils_logus.FilePath(utils_types.FilePath(settings.FreelancerFreelancerLocation)))
	}

	configs := configs_mapped.NewMappedConfigs().Read(utils_types.FilePath(settings.FreelancerFreelancerLocation))

	// see README.go in denormalizer why it was commented out but not removed.
	// denormalizer.Run(data)

	formator := NewFormatter(configs)
	formator.Format()

	configs.Write(is_dry_run)
}
