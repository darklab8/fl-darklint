/*
Scanned info with package `parser` we try here to validate for being correct
And even suggesting autofixes to Freelancer config files
*/
package validator

import (
	"darktool/settings"
	"darktool/tools/parser"
)

func Run() {

	parser.Run(settings.DryRun)
}
