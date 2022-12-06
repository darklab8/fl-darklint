/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/settings"
	"darktool/tools/parser/freelancer/data/equipment/market"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/data/universe/systems"
	"darktool/tools/parser/freelancer/infocard"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	log "github.com/sirupsen/logrus"
)

var folderToLowerCaseRegex *regexp.Regexp

func initRegexExpression(regex **regexp.Regexp, expression string) {
	var err error

	*regex, err = regexp.Compile(expression)
	utils.CheckFatal(err, "failed to parse numberParser in ", utils.GetCurrentFile())
}

func init() {
	initRegexExpression(&folderToLowerCaseRegex, `universe(\/|\\)`)
}
func Parse(file1path string, dry_run bool) {
	log.Info("Parse START for FreelancerFolderLocation=", file1path)
	filesystem := filefind.FindConfigs(file1path)

	universe_config := universe.Config{}
	universe_config.Read(filesystem.GetFile(universe.FILENAME))
	universe_config.Write().WriteLines(dry_run)
	info_config := infocard.Config{}
	info_config.Read(filesystem.GetFile(infocard.FILENAME, infocard.FILENAME_FALLBACK))

	systems := (&systems.Config{}).Read(&universe_config, filesystem)
	system_files := systems.Write()
	for _, system_file := range system_files {
		system_file.WriteLines(dry_run)
	}

	market_ships_config := market.Config{}
	market_ships_config.Read(filesystem.GetFile(market.FILENAME_SHIPS))
	market_ships_config.UpdateWithBasenames(&universe_config, &info_config)
	market_ships_config.UpdateWithRecycle(&universe_config, systems)
	market_ships_config.Write().WriteLines(dry_run)

	market_commodities := market.Config{}
	market_commodities.Read(filesystem.GetFile(market.FILENAME_COMMODITIES))
	market_commodities.UpdateWithBasenames(&universe_config, &info_config)
	market_commodities.UpdateWithRecycle(&universe_config, systems)
	market_commodities.Write().WriteLines(dry_run)

	market_misc := market.Config{}
	market_misc.Read(filesystem.GetFile(market.FILENAME_MISC))
	market_misc.UpdateWithBasenames(&universe_config, &info_config)
	market_misc.Write().WriteLines(dry_run)

	log.Info("Parse OK for FreelancerFolderLocation=", file1path)
}

func Run(dry_run bool) {
	_, err := os.Stat(filepath.Join(settings.FreelancerFreelancerLocation, "DATA"))

	fmt.Println(err)
	if os.IsNotExist(err) {
		log.Fatal("freelancer folder is not detected at path=", settings.FreelancerFreelancerLocation, " because DATA folder was not found")
	}

	Parse(settings.FreelancerFreelancerLocation, dry_run)
}
