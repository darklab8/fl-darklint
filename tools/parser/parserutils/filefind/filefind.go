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

type FileInfo struct {
	AbsPath string
}

type Filesystem struct {
	Files   []FileInfo
	Hashmap map[string]FileInfo
}

var FreelancerFolder Filesystem

func FindConfigs(folderpath string) Filesystem {
	var filesystem Filesystem
	filesystem.Hashmap = make(map[string]FileInfo)

	err := filepath.WalkDir(folderpath, func(path string, d fs.DirEntry, err error) error {

		if !strings.Contains(path, ".ini") {
			return nil
		}

		utils.CheckFatal(err, "unable to read file")

		file := FileInfo{AbsPath: path}
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
