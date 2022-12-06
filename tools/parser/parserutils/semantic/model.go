package semantic

import (
	"darktool/tools/parser/parserutils/inireader"
	"darktool/tools/utils"
)

type Model struct {
	section *inireader.Section
}

func (s *Model) Map(section *inireader.Section) {
	s.section = section
}

func (s *Model) Render() *inireader.Section {
	return s.section
}

type ConfigModel struct {
	sections []*inireader.Section
	comments []string
	filepath string
}

func (s *ConfigModel) Init(sections []*inireader.Section, comments []string, filepath string) {
	s.sections = sections
	s.comments = comments
	s.filepath = filepath
}

func (s *ConfigModel) SetOutputPath(filepath string) {
	s.filepath = filepath
}

func (s *ConfigModel) Render() *inireader.INIFile {
	inifile := &inireader.INIFile{}
	inifile.Comments = s.comments
	inifile.Sections = s.sections
	inifile.File = &utils.File{Filepath: s.filepath}
	return inifile
}
