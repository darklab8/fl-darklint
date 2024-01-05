/*
Scanned info with package `parser` we try here to validate for being correct
And even suggesting autofixes to Freelancer config files
*/
package validator

import (
	"darklint/fldarklint/settings"
	"darklint/fldarklint/settings/logus"
	"fmt"
	"os"
	"path/filepath"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/configs_mapped"
	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func Run(is_dry_run configs_mapped.IsDruRun) {

	_, err := os.Stat(filepath.Join(settings.FreelancerFreelancerLocation, "DATA"))

	fmt.Println(err)
	if os.IsNotExist(err) {
		logus.Log.Fatal("freelancer folder is not detected because DATA folder was not found", logus_core.FilePath(utils_types.FilePath(settings.FreelancerFreelancerLocation)))
	}

	configs := configs_mapped.NewMappedConfigs().Read(utils_types.FilePath(settings.FreelancerFreelancerLocation))

	// see README.go in denormalizer why it was commented out but not removed.
	// denormalizer.Run(data)

	configs.Write(is_dry_run)
}
