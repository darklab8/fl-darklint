/*
Okay we need to create syntax. To augment currently possible stuff
*/
package inireader

import (
	"darktool/tools/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type INIFile struct {
	File     *utils.File
	Comments []string

	Sections []*Section

	// denormalization
	SectionMap map[string][]*Section

	// Enforce unique keys
	ConstraintUniqueSectionType map[string]string
}

func (config *INIFile) AddSection(key string, section *Section) {
	config.Sections = append(config.Sections, section)

	// Denormalization adding to hashmap
	if config.SectionMap == nil {
		config.SectionMap = make(map[string][]*Section)
	}
	if _, ok := config.SectionMap[key]; !ok {
		config.SectionMap[key] = make([]*Section, 0)
	}
	config.SectionMap[key] = append(config.SectionMap[key], section)

	// Enforcing same case sensetivity for section type key
	if config.ConstraintUniqueSectionType == nil {
		config.ConstraintUniqueSectionType = make(map[string]string)
	}

	if val, ok := config.ConstraintUniqueSectionType[strings.ToLower(key)]; ok {
		if val != key {
			log.Fatal("not uniform case sensetivity for config.path", config.File.Filepath, " key=", key, " section=", section)
		}
	} else {
		config.ConstraintUniqueSectionType[strings.ToLower(key)] = key
	}
}

/*
[BaseGood] // this is Type
abc = 123 // this is Param going into list and hashmap
*/
type Section struct {
	Type   string
	Params []*Param
	// denormialization of Param list due to being more comfortable
	ParamMap map[string][]*Param
}

const (
	OPTIONAL_p = true
	REQUIRED_p = false
)

func (section *Section) AddParam(key string, param *Param) {
	param.Key = key

	section.Params = append(section.Params, param)
	// Denormalization, adding to hashmap
	if section.ParamMap == nil {
		section.ParamMap = make(map[string][]*Param)
	}
	if _, ok := section.ParamMap[key]; !ok {
		section.ParamMap[key] = make([]*Param, 0)
	}
	section.ParamMap[key] = append(section.ParamMap[key], param)
}

func (section *Section) GetParamStr(key string, optional bool) string {
	if optional && len(section.ParamMap[key]) == 0 {
		return ""
	}
	return section.ParamMap[key][0].First.AsString()
}
func (section *Section) GetParamStrToLower(key string, optional bool) string {
	return strings.ToLower(section.GetParamStr(key, optional))
}
func (section *Section) GetParamInt(key string, optional bool) int {
	if optional && len(section.ParamMap[key]) == 0 {
		return 0
	}

	integer, err := strconv.Atoi(section.GetParamStr(key, false))
	if err != nil {
		log.Fatal("failed to parse strid in universe.ini for section=", section, "key=", key)
	}
	return integer
}
func (section *Section) GetParamNumber(key string, optional bool) ValueNumber {
	if optional && len(section.ParamMap[key]) == 0 {
		return ValueNumber{}
	}

	return section.ParamMap[key][0].First.(ValueNumber)
}
func (section *Section) GetParamBool(key string, optional bool, default_value bool) bool {
	if optional && len(section.ParamMap[key]) == 0 {
		return default_value
	}

	bool_value, err := strconv.ParseBool(section.GetParamStr(key, REQUIRED_p))

	if err != nil {
		return default_value
	}

	return bool_value
}

// abc = qwe, 1, 2, 3, 4
// abc is key
// qwe is first value
// qwe, 1, 2, 3, 4 are values
// ;abc = qwe, 1, 2, 3 is Comment
type Param struct {
	Key       string
	Values    []UniValue
	IsComment bool     // if commented out
	First     UniValue // denormalization due to very often being needed
}

func (p *Param) AddValue(value UniValue) *Param {
	if len(p.Values) == 0 {
		p.First = value
	}
	p.Values = append(p.Values, value)
	return p
}

func (p Param) ToString() string {
	var sb strings.Builder

	if p.IsComment {
		sb.WriteString(";%")
	}

	sb.WriteString(fmt.Sprintf("%v = ", p.Key))

	for index, value := range p.Values {
		str_to_write := value.AsString()
		if index == len(p.Values)-1 {
			sb.WriteString(str_to_write)
		} else {
			sb.WriteString(fmt.Sprintf("%v, ", str_to_write))
		}
	}

	return sb.String()
}

type UniValue interface {
	AsString() string
}
type ValueString string
type ValueNumber struct {
	Value     float64
	Precision int
}

type ValueBool bool

func (v ValueBool) AsString() string {
	return strconv.FormatBool(bool(v))
}

func (v ValueString) AsString() string {
	return string(v)
}

func (v ValueString) ToLowerValue() ValueString {
	return ValueString(strings.ToLower(string(v)))
}

func (v ValueNumber) AsString() string {
	return strconv.FormatFloat(float64(v.Value), 'f', v.Precision, 64)
}

func UniParse(input string) (UniValue, error) {

	input = strings.ReplaceAll(input, " ", "")

	numberMatch := regexNumber.FindAllString(input, -1)
	if len(numberMatch) > 0 && !strings.ContainsAny(input, "") {
		parsed_number, err := strconv.ParseFloat(input, 64)

		if err != nil {
			log.Warn("failed to read number, input=", input)
			return nil, err
		}

		var precision int
		if !strings.Contains(input, ".") {
			precision = 0
		} else {
			precision = 1
		}

		return ValueNumber{Value: parsed_number, Precision: precision}, nil
	}

	v := ValueString(input)
	return v, nil
}
func UniParseF(input string) UniValue {
	value, err := UniParse(input)
	if err != nil {
		log.Fatal("unable to parse UniParseF=", input)
	}
	return value
}

func UniParseStr(input string) UniValue {
	return ValueString(input)
}

func UniParseInt(input int) UniValue {
	u := ValueNumber{}
	u.Value = float64(input)
	u.Precision = 0
	return u
}

var regexNumber *regexp.Regexp
var regexComment *regexp.Regexp
var regexSection *regexp.Regexp
var regexParam *regexp.Regexp

func init() {
	initRegexExpression(&regexNumber, `^[0-9\-]+(?:\.)?([0-9\-]*)$`)
	initRegexExpression(&regexComment, `;(.*)`)
	initRegexExpression(&regexSection, `^\[.*\]`)
	// param or commented out param
	initRegexExpression(&regexParam, `(;%|^)([a-zA-Z_][a-zA-Z_0-9]+)\s=\s([a-zA-Z_, 0-9-.\\]+)`)
}

var CASE_SENSETIVE_KEYS = [...]string{"BGCS_base_run_by", "NavMapScale"}

func isKeyCaseSensetive(key string) bool {
	for _, sensetive_key := range CASE_SENSETIVE_KEYS {
		if key == sensetive_key {
			return true
		}
	}
	return false
}

func (config INIFile) Read(fileref *utils.File) INIFile {
	log.Debug("started reading INIFileRead for", fileref.Filepath)
	config.File = fileref

	log.Debug("opening file", fileref.Filepath)
	file := fileref.OpenToReadF()
	log.Debug("defer file close", fileref.Filepath)
	defer file.Close()

	log.Debug("reading lines")
	lines := file.ReadLines()

	log.Debug("setting current section")
	var cur_section *Section
	cur_section = &Section{}
	for _, line := range lines {

		comment_match := regexComment.FindStringSubmatch(line)
		section_match := regexSection.FindStringSubmatch(line)
		param_match := regexParam.FindStringSubmatch(line)

		if len(param_match) > 0 {
			isComment := len(param_match[1]) > 0
			key := param_match[2]
			if !isKeyCaseSensetive(key) {
				key = strings.ToLower(key)
			}

			line_to_read := param_match[3]
			if strings.Contains(line_to_read, ",") {
				line_to_read = strings.ReplaceAll(line_to_read, " ", "")
			}
			splitted_values := strings.Split(line_to_read, ",")
			first_value, err := UniParse(splitted_values[0])
			if err != nil {
				log.Fatal("ini reader, failing to parse line because of UniParse, line=", line, "filepath=", fileref.Filepath)
			}

			var values []UniValue
			for _, value := range splitted_values {
				univalue, err := UniParse(value)
				if err != nil {
					log.Fatal("ini reader, failing to parse line because of UniParse, line=", line, "filepath=", fileref.Filepath)
				}
				values = append(values, univalue)
			}

			param := Param{Key: key, First: first_value, Values: values, IsComment: isComment}
			cur_section.AddParam(key, &param)
		} else if len(comment_match) > 0 {
			config.Comments = append(config.Comments, comment_match[1])
		} else if len(section_match) > 0 {
			cur_section = &Section{} // create new
			cur_section.Type = section_match[0]
			config.AddSection(section_match[0], cur_section)
		}

	}

	return config
}

func (config INIFile) Write(fileref *utils.File) *utils.File {

	for _, comment := range config.Comments {
		fileref.ScheduleToWrite(fmt.Sprintf(";%s", comment))
	}

	if len(config.Comments) > 0 {
		fileref.ScheduleToWrite("")
	}

	section_length := config.Sections
	for index, section := range config.Sections {
		fileref.ScheduleToWrite(section.Type)

		for _, param := range section.Params {
			fileref.ScheduleToWrite(param.ToString())
		}

		if index < len(section_length)-1 {
			fileref.ScheduleToWrite("")
		}
	}

	return fileref
}
