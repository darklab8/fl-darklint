/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/tools/parser/data/equipment/market1ships"
	"darktool/tools/parser/parserutils/filefind"
)

func Run() {
	filefind.LoadFreelancerConfigs()
	market1ships.Parse()
}