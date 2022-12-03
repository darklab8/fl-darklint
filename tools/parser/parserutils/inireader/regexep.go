package inireader

import (
	"darktool/tools/utils"
	"regexp"
)

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
