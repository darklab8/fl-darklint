/*
ORM mapper for Freelancer ini reader
*/
package semantic

import "darktool/tools/parser/parserutils/inireader"

// ORM values

const (
	TypeComment = true
	TypeVisible = false
)

type Value struct {
	section   *inireader.Section
	key       string
	optional  bool
	isComment bool
}

type String struct {
	Value
}

func (s *String) Map(section *inireader.Section, key string, isComment bool, optional bool) *String {
	s.section = section
	s.key = key
	s.optional = optional
	s.isComment = isComment
	return s
}

func (s *String) Get() string {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return ""
	}
	return s.section.ParamMap[s.key][0].First.AsString()
}

func (s *String) Set(value string) {
	if s.isComment {
		s.Delete()
	}

	processed_value := inireader.UniParseStr(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment}).AddValue(processed_value))
	}
	// implement SetValue in Section
	s.section.ParamMap[s.key][0].First = processed_value
	s.section.ParamMap[s.key][0].Values[0] = processed_value
}

func (s *String) Delete() {
	delete(s.section.ParamMap, s.key)
	for index, param := range s.section.Params {
		if param.Key == s.key {
			s.section.Params = append(s.section.Params[:index], s.section.Params[index+1:]...)
		}
	}
}

// ORM Model
