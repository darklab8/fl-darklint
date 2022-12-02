/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/tools/parser/data/equipment/market1ships"
	"darktool/tools/parser/file1discovery"
)

func Run() {
	file1discovery.DiscoverFiles()
	market1ships.Parse()
}
