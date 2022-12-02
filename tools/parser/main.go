/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/tools/parser/data/equipment/market1ships"
	"darktool/tools/parser/parserutils"
)

func Run() {
	parserutils.DiscoverFiles()
	market1ships.Parse()
}
