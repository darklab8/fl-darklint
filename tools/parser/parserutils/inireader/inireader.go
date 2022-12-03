/*
Okay we need to create syntax. To augment currently possible stuff
*/
package inireader

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type INIFile struct {
	File     filefind.FileInfo
	Sections []*Section
	Comments []string
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

		return ValueNumber{value: parsed_number, precision: len(numberMatch[1])}
	}

	return ValueString(input)
}

var regexNumber *regexp.Regexp
var regexComment *regexp.Regexp
var regexSection *regexp.Regexp
var regexParam *regexp.Regexp

func init() {
	initRegexExpression(&regexNumber, `[0-9\-]+(?:\.)?([0-9\-]*)`)
	initRegexExpression(&regexComment, `;(.*)`)
	initRegexExpression(&regexSection, `^\[.*\]`)
	initRegexExpression(&regexParam, `^([a-zA-Z_]+)\s=\s([a-zA-Z_, 0-9-]+)`)
}

func INIFileRead(file1path string) INIFile {
	log.Debug("started reading INIFileRead for", file1path)
	config := INIFile{}

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
		// param_match := regexParam.FindStringSubmatch(line)

		if len(comment_match) > 0 {
			config.Comments = append(config.Comments, comment_match[1])
		} else if len(section_match) > 0 {
			cur_section = &Section{}
			config.Sections = append(config.Sections, cur_section)
			continue
		}

		// 	if len(param_match) > 0 {
		// 		if strings.Compare(param_match[1], "base") == 0 {
		// 			current_base_good.Base = param_match[2]
		// 		} else if strings.Compare(param_match[1], "MarketGood") == 0 {
		// 			params := strings.Split(param_match[2], ", ") // data example: dsy_arrow_package, 1, -1, 1, 1, 0, 1, 1
		// 			var floats []float32

		// 			for _, string_number := range params {
		// 				parsed_float, _ := strconv.ParseFloat(string_number, 32)
		// 				floats = append(floats, float32(parsed_float))
		// 			}

		// 			current_base_good.Goods = append(current_base_good.Goods, MarketGood{Name: params[0], Values: floats})
		// 		}
		// 		continue
		// 	}
	}

	return config
}

// // comments
// comment_exp, err := regexp.Compile(`;(.*)`)
// utils.CheckPanic(err)
// // [BaseGood]
// base_group_ext, err := regexp.Compile(`^\[.*\]`)
// utils.CheckPanic(err)
// // `base = br01_01_base` or `MarketGood = dsy_arrow_package, 1, -1, 1, 1, 0, 1, 1`
// param_exp, err := regexp.Compile(`^([a-zA-Z_]+)\s=\s([a-zA-Z_, 0-9-]+)`)
// utils.CheckPanic(err)
