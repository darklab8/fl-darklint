package denormalizer

import "github.com/darklab8/fl-configs/configs/configs_mapped"

/*
This module is able to add useful human readable comments to config objects
regarding what they are. For example human readable base names to base object definitions
*/

func Run(configs *configs_mapped.MappedConfigs) {
	(&BaseDenormalizer{}).Read(configs).Write(configs)
}
