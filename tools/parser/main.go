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
	// log.Info("filefind. Checking fo", file1path)
	// filesystem := filefind.FindConfigs(file1path)
	// log.Info("")
	// for _, file := range filesystem.Files {
	// 	if strings.Contains(strings.ToLower(file.Filepath), "universe") && !utils.IsLower(filepath.Base(file.Filepath)) {
	// 		os.Rename(file.Filepath, strings.ToLower(file.Filepath))
	// 	}
	// }

	log.Info("Parse START for FreelancerFolderLocation=", file1path)
	filesystem := filefind.FindConfigs(file1path)

	universe_config := universe.Config{}
	universe_config.Read(filesystem.GetFile(universe.FILENAME))
	universe_config.Write(filesystem.GetFile(universe.FILENAME)).WriteLines(dry_run)
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

	for _, base := range universe_config.Bases {
		end_objects := universe_config.SystemMap[universe.SystemNickname(base.System)].File.PathObjects()
		for i := 0; i < 2; i++ {
			end_objects[i] = strings.ToUpper(end_objects[i])
		}
		objects := append([]string{file1path, "DATA", "UNIVERSE"}, end_objects...)
		path_to_systemfile := filepath.Join(objects...)
		fmt.Println(path_to_systemfile)
	}

	// TODO implement preserving comments :|
	// market_misc := market1ships.Config{}
	// market_misc.Read(filesystem.GetFile(market1ships.FILENAME_MISC))
	// market_misc.UpdateWithBasenames(&universe_config, &info_config)
	// market_misc.Write(filesystem.GetFile(market1ships.FILENAME_MISC)).WriteLines(dry_run)

	log.Info("Parse OK for FreelancerFolderLocation=", file1path)
}

func Run(dry_run bool) {
	_, err := os.Stat(filepath.Join(settings.FreelancerFreelancerLocation, "DATA"))
	_, err2 := os.Stat(filepath.Join(settings.FreelancerFreelancerLocation, "SERVICE"))

	fmt.Println(err)
	fmt.Println(err2)
	if os.IsNotExist(err) || os.IsNotExist(err2) {
		log.Fatal("freelancer folder is not detected at path=", settings.FreelancerFreelancerLocation, " because not found DATA and SERVICE folders")
	}

	Parse(settings.FreelancerFreelancerLocation, dry_run)
}
