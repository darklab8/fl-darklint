/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/settings"
	"darktool/tools/parser/freelancer/data/equipment/market"
	"darktool/tools/parser/freelancer/data/universe"
	"darktool/tools/parser/freelancer/service"
	"darktool/tools/parser/parserutils/filefind"
	"darktool/tools/utils"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
	log.Info("filefind. Checking fo", file1path)
	filesystem := filefind.FindConfigs(file1path)
	log.Info("")
	for _, file := range filesystem.Files {
		if strings.Contains(strings.ToLower(file.Filepath), "universe") {
			fmt.Println("debug")
		}

		if strings.Contains(strings.ToLower(file.Filepath), "universe") && !utils.IsLower(filepath.Base(file.Filepath)) {
			os.Rename(file.Filepath, strings.ToLower(file.Filepath))
		}
	}

	log.Info("Parse START for FreelancerFolderLocation=", file1path)
	filesystem = filefind.FindConfigs(file1path)

	universe_config := universe.Config{}
	_, iniconfig := universe_config.Read(filesystem.GetFile(universe.FILENAME))
	iniconfig.Write(filesystem.GetFile(universe.FILENAME)).WriteLines(dry_run)

	info_config := service.Config{}
	info_config.Read(filesystem.GetFile(service.FILENAME, service.FILENAME_FALLBACK))

	market_ships_config := market.Config{}
	market_ships_config.Read(filesystem.GetFile(market.FILENAME_SHIPS))
	market_ships_config.UpdateWithBasenames(&universe_config, &info_config)
	market_ships_config.Write(filesystem.GetFile(market.FILENAME_SHIPS)).WriteLines(dry_run)

	market_commodities := market.Config{}
	market_commodities.Read(filesystem.GetFile(market.FILENAME_COMMODITIES))
	market_commodities.UpdateWithBasenames(&universe_config, &info_config)
	market_commodities.Write(filesystem.GetFile(market.FILENAME_COMMODITIES)).WriteLines(dry_run)

	// TODO implement preserving comments :|
	// market_misc := market1ships.Config{}
	// market_misc.Read(filesystem.GetFile(market1ships.FILENAME_MISC))
	// market_misc.UpdateWithBasenames(&universe_config, &info_config)
	// market_misc.Write(filesystem.GetFile(market1ships.FILENAME_MISC)).WriteLines(dry_run)

	log.Info("Parse OK for FreelancerFolderLocation=", file1path)
}

func Run(dry_run bool) {
	Parse(settings.FreelancerFolderLocation, dry_run)
}
