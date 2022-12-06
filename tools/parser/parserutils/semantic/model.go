package semantic

import "darktool/tools/parser/parserutils/inireader"

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
	Sections []*inireader.Section
	Comments []string
	Filepath string
}

func (s *ConfigModel) Init(sections []*inireader.Section, comments []string, filepath string) {
	s.Sections = sections
	s.Comments = comments
	s.Filepath = filepath
}
