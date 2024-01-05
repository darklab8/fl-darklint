package formatter

import (
	"darklint/fldarklint/formatter/freelancer_format/data_format/universe_format"
	"darklint/fldarklint/settings"
	"darklint/fldarklint/settings/logus"
	"fmt"
	"os"
	"path/filepath"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped"
	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
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
	f.formatters = append(f.formatters, universe_format.NewUniverseFormatter(configs.Universe_config))
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
		logus.Log.Fatal("freelancer folder is not detected because DATA folder was not found", logus_core.FilePath(utils_types.FilePath(settings.FreelancerFreelancerLocation)))
	}

	configs := configs_mapped.NewMappedConfigs().Read(utils_types.FilePath(settings.FreelancerFreelancerLocation))

	// see README.go in denormalizer why it was commented out but not removed.
	// denormalizer.Run(data)

	formator := NewFormatter(configs)
	formator.Format()

	configs.Write(is_dry_run)
}
