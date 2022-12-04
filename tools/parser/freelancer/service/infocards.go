package service

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Record interface {
	Id() int
	Type() string
	Content() string
}
type Infocard struct {
	id      int
	content string
}
type Name struct {
	id      int
	content string
}

func (v Name) Id() int {
	return v.id
}
func (v Infocard) Id() int {
	return v.id
}
func (v Name) Content() string {
	return v.content
}
func (v Infocard) Content() string {
	return v.content
}

func (v Name) Type() string {
	return "NAME"
}
func (v Infocard) Type() string {
	return "INFOCARD"
}

type Config struct {
	Records []*Record

	RecordsMap map[int]*Record
}

var LoadedInfocards *Config

const (
	Filename = "infocards.txt"
)

func (frelconfig *Config) Read(input_file *utils.File) *Config {
	frelconfig.RecordsMap = make(map[int]*Record)
	frelconfig.Records = make([]*Record, 0)

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

		var record_to_add Record
		switch name {
		case "NAME":
			record_to_add = Name{id: id, content: content}
		case "INFOCARD":
			record_to_add = Infocard{id: id, content: content}
		default:
			log.Fatal("unrecognized object name in infocards.txt, id=", id, "name=", name, "content=", content)
		}

		frelconfig.Records = append(frelconfig.Records, &record_to_add)
		frelconfig.RecordsMap[record_to_add.Id()] = &record_to_add
	}

	return frelconfig
}

func Load() {
	file := &utils.File{Filepath: filefind.FreelancerFolder.Hashmap[Filename].Filepath}
	config := Config{}
	LoadedInfocards = config.Read(file)
	log.Info("OK ", Filename, " is parsed to specialized data structs")
}
