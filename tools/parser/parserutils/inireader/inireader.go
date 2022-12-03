/*
Okay we need to create syntax. To augment currently possible stuff
*/
package inireader

import (
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"regexp"
	"strconv"
)

type INIFile struct {
	File     filefind.FileInfo
	Sections []Section
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

	numberMatch := numberParser.FindAllString(input, -1)
	if len(numberMatch) > 0 {
		parsed_number, err := strconv.ParseFloat(input, 64)
		utils.CheckFatal(err, "failed to read number, input=", input)

		return ValueNumber{value: parsed_number, precision: len(numberMatch[1])}
	}

	return ValueString(input)
}

var numberParser *regexp.Regexp

func initNumberParser() {
	regex_exp := `[0-9\-]+(?:\.)?([0-9\-]*)`
	var err error
	numberParser, err = regexp.Compile(regex_exp)
	utils.CheckFatal(err, "failed to parse regex_exp=", regex_exp)
}

func init() {
	initNumberParser()
}
