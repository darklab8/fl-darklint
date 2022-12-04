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

	log "github.com/sirupsen/logrus"
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

		if !strings.Contains(path, ".ini") && !strings.Contains(path, ".txt") && !strings.Contains(path, ".xml") {
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

func Load() {
	if len(FreelancerFolder.Files) == 0 {
		FreelancerFolder = FindConfigs(settings.FreelancerFolderLocation)
	}
}

func (file1system Filesystem) GetFile(file1names ...string) *utils.File {
	for _, file1name := range file1names {
		file, ok := file1system.Hashmap[file1name]
		if !ok {
			log.Warn("Filesystem.GetFile, failed to find find in filesystesm filename=", file1name, ", trying to recover")
			continue
		}
		log.Info("Filesystem.GetFile, found filepath=", file.Filepath)
		result_file := utils.File{Filepath: file.Filepath}
		return &result_file
	}
	log.Fatal("unable to find filenames=", file1names)
	return nil
}
