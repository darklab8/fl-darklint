package service

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Record interface {
}
type Infocard struct {
	id      int
	content string
}

type Name struct {
	id      int
	content string
}

type Config struct {
	Records []Record
}

var Loaded Config

const (
	filename = "infocards.txt"
)

func Read(input_file utils.File) Config {
	var frelconfig Config

	input_file = input_file.OpenToReadF()
	defer input_file.Close()
	lines := input_file.ReadLines()

	for index := 0; index < len(lines); index++ {
		id, err := strconv.Atoi(lines[index])
		if err != nil {
			continue
		}
		name := lines[index+1]
		content := lines[index+2]
		index += 3

		switch name {
		case "NAME":
			frelconfig.Records = append(frelconfig.Records, Name{id: id, content: content})
		case "INFOCARD":
			frelconfig.Records = append(frelconfig.Records, Name{id: id, content: content})
		default:
			log.Fatal("unrecognized object name in infocards.txt, id=", id, "name=", name, "content=", content)
		}

	}

	return frelconfig
}

func Load() {
	file := utils.File{Filepath: filefind.FreelancerFolder.Hashmap[filename].Filepath}
	Read(file)
	log.Info("OK ", filename, " is parsed to specialized data structs")
}
