/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/tools/parser/freelancer/data/equipment/market1ships"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"
	"darktool/tools/parser/parserutils/filefind"
)

func Run() {
	filefind.Load()
	market1ships.Load()
	universe.Load()
	service.Load()
}
