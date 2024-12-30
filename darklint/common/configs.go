package common

import (
	"fmt"
	"os"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darklint/darklint/settings/logus"
	"github.com/darklab8/go-utils/utils/utils_logus"
	"github.com/darklab8/go-utils/utils/utils_types"
)

func GetConfigs(freelancer_folder utils_types.FilePath) *configs_mapped.MappedConfigs {
	_, err := os.Stat(freelancer_folder.Join("DATA").ToString())

	fmt.Println(err)
	if os.IsNotExist(err) {
		logus.Log.Fatal("freelancer folder is not detected because DATA folder was not found", utils_logus.FilePath(freelancer_folder))
	}

	configs := configs_mapped.NewMappedConfigs().Read(utils_types.FilePath(freelancer_folder))
	return configs
}
