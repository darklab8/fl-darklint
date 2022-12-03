/*
Package with reusable code for discovery of files and other reusable stuff like universal ini reader
*/
package filefind

import (
	"darktool/settings"
	"darktool/tools/utils"
	"io/fs"
	"path/filepath"
	"strings"
)

type Filesystem struct {
	Files   []utils.File
	Hashmap map[string]utils.File
}

var FreelancerFolder Filesystem

func FindConfigs(folderpath string) Filesystem {
	var filesystem Filesystem
	filesystem.Hashmap = make(map[string]utils.File)

	err := filepath.WalkDir(folderpath, func(path string, d fs.DirEntry, err error) error {

		if !strings.Contains(path, ".ini") {
			return nil
		}

		utils.CheckFatal(err, "unable to read file")

		file := utils.File{Filepath: path}
		filesystem.Files = append(filesystem.Files, file)

		key := filepath.Base(path)
		filesystem.Hashmap[key] = file

		return nil
	})

	utils.CheckFatal(err, "unable to read files")
	return filesystem
}

func LoadFreelancerConfigs() {
	FreelancerFolder = FindConfigs(settings.FreelancerFolderLocation)
}
