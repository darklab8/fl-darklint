package formatter

import (
	"strings"
	"sync"

	"github.com/darklab8/fl-darklint/darklint/denormalizer"
	"github.com/darklab8/fl-darklint/darklint/formatter/freelancer_format/data_format/equipment_format/market_format"
	"github.com/darklab8/fl-darklint/darklint/formatter/freelancer_format/data_format/universe_format"
	"github.com/darklab8/fl-darklint/darklint/formatter/freelancer_format/data_format/universe_format/systems_mapped"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/go-utils/utils/utils_types"
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
	f.formatters = append(f.formatters, universe_format.NewFormatter(configs.Universe))
	f.formatters = append(f.formatters, systems_mapped.NewFormatter(configs.Systems))

	f.formatters = append(f.formatters,
		market_format.NewFormatter(f.configs.Market),
	)
	return f
}

func (f *Formatter) Format() {
	for _, formatter := range f.formatters {
		formatter.Format()
	}
}

func Run(configs *configs_mapped.MappedConfigs, freelancer_folder utils_types.FilePath, is_dry_run configs_mapped.IsDruRun) {

	denormalizer.Run(configs)

	formator := NewFormatter(configs)
	formator.Format()

	configs.Write(is_dry_run)

	if !is_dry_run {
		ReformatAll(freelancer_folder)
	}
}

func ReformatAll(FreelancerFolder utils_types.FilePath) {
	filesystem := filefind.FindConfigs(FreelancerFolder)

	var ini_files []*iniload.IniLoader
	for filepath, file := range filesystem.Hashmap {
		if file.IsFailback {
			continue
		}

		if strings.Contains(filepath.Base().ToString(), "ini") {
			ini_files = append(ini_files, iniload.NewLoader(file))
		}
	}

	var wg sync.WaitGroup
	for _, file := range ini_files {
		wg.Add(1)
		go func(file *iniload.IniLoader) {
			file.Scan()
			file.File.WriteLines()
			wg.Done()
		}(file)
	}
	wg.Wait()
}
