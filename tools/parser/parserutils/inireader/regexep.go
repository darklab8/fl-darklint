package inireader

import (
	"darktool/tools/utils"
	"regexp"
)

func initRegexExpression(regex **regexp.Regexp, expression string) {
	var err error

	*regex, err = regexp.Compile(expression)
	utils.CheckFatal(err, "failed to parse numberParser in ", utils.GetCurrentFile())
}
