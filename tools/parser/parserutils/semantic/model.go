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
