package cmd

import (
	"github.com/darklab8/fl-darklint/darklint/settings"
	"github.com/darklab8/go-utils/utils/utils_types"
)

var is_dry_run bool
var freelancer_folder string

func GetFreelancerFolder() utils_types.FilePath {
	var FreelancerFolderTarget utils_types.FilePath
	if freelancer_folder != "" {
		FreelancerFolderTarget = utils_types.FilePath(freelancer_folder)
	} else {
		FreelancerFolderTarget = settings.Env.FreelancerFolder
	}
	return FreelancerFolderTarget
}
