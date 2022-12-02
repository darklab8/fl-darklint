/*
Package with reusable code for discovery of files and other reusable stuff like universal ini reader
*/
package parserutils

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

var Filesystem struct {
	Files   []FileInfo
	Hashmap map[string]FileInfo
}

func DiscoverFiles() {
	Filesystem.Hashmap = make(map[string]FileInfo)

	err := filepath.WalkDir(settings.FreelancerFolderLocation, func(path string, d fs.DirEntry, err error) error {

		if !strings.Contains(path, ".ini") {
			return nil
		}

		utils.CheckFatal(err, "unable to read file")

		file := FileInfo{AbsPath: path}
		Filesystem.Files = append(Filesystem.Files, file)

		key := filepath.Base(path)
		Filesystem.Hashmap[key] = file

		return nil
	})

	utils.CheckFatal(err, "unable to read files")
}
