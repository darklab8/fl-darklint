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
	File     utils.File
	Comments []string

	Sections []*Section
}

/*
[BaseGood] // this is Type
abc = 123 // this is Param going into list and hashmap
*/
type Section struct {
	Type   string
	Params []Param
	// denormialization of Param list due to being more comfortable
	ParamMap map[string][]Param
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

type UniValue interface {
	AsString() string
}
type ValueString string
type ValueNumber struct {
	value     float64
	precision int
}

func (v ValueString) AsString() string {
	return string(v)
}
func (v ValueNumber) AsString() string {
	return strconv.FormatFloat(float64(v.value), 'f', v.precision, 64)
}

func UniParse(input string) UniValue {

	numberMatch := regexNumber.FindAllString(input, -1)
	if len(numberMatch) > 0 {
		parsed_number, err := strconv.ParseFloat(input, 64)
		utils.CheckFatal(err, "failed to read number, input=", input)

		var precision int
		if len(numberMatch) == 1 {
			precision = 0
		} else {
			precision = len(numberMatch[1])
		}

		return ValueNumber{value: parsed_number, precision: precision}
	}

	v := ValueString(input)
	return v
}

var regexNumber *regexp.Regexp
var regexComment *regexp.Regexp
var regexSection *regexp.Regexp
var regexParam *regexp.Regexp

func init() {
	initRegexExpression(&regexNumber, `^[0-9\-]+(?:\.)?([0-9\-]*)`)
	initRegexExpression(&regexComment, `;(.*)`)
	initRegexExpression(&regexSection, `^\[.*\]`)
	initRegexExpression(&regexParam, `^([a-zA-Z_]+)\s=\s([a-zA-Z_, 0-9-]+)`)
}

func INIFileRead(file1path string) INIFile {
	log.Debug("started reading INIFileRead for", file1path)
	config := INIFile{File: utils.File{Filepath: file1path}}

	log.Debug("opening file", file1path)
	file := utils.File.OpenToReadF(utils.File{Filepath: file1path})
	log.Debug("defer file close", file1path)
	defer file.Close()

	log.Debug("reading lines")
	lines := file.ReadLines()

	log.Debug("setting current section")
	var cur_section *Section = &Section{}
	for _, line := range lines {

		log.Debug("reading regex")
		comment_match := regexComment.FindStringSubmatch(line)
		section_match := regexSection.FindStringSubmatch(line)
		param_match := regexParam.FindStringSubmatch(line)

		if len(comment_match) > 0 {
			config.Comments = append(config.Comments, comment_match[1])
		} else if len(section_match) > 0 {
			cur_section := Section{}
			config.Sections = append(config.Sections, &cur_section)
			cur_section.Type = section_match[0]
		} else if len(param_match) > 0 {
			fmt.Println("123")
			key := param_match[1]
			splitted_values := strings.Split(param_match[2], ", ")
			first_value := UniParse(splitted_values[0])
			var values []UniValue
			for _, value := range splitted_values {
				values = append(values, UniParse(value))
			}

			// TODO add reading commented param
			param := Param{Key: key, First: first_value, Values: values, IsComment: false}
			cur_section.Params = append(cur_section.Params, param)

			// Denormalization
			if cur_section.ParamMap == nil {
				cur_section.ParamMap = make(map[string][]Param)
			}
			if _, ok := cur_section.ParamMap[key]; !ok {
				cur_section.ParamMap[key] = make([]Param, 0)
			}
			cur_section.ParamMap[key] = append(cur_section.ParamMap[key], param)
		}

	}

	return config
}
