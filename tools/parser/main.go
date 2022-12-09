/*
Tool to parse freelancer configs
*/
package parser

import (
	"darktool/settings"
	"darktool/tools/denormalizer"
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

	// read
	universe_config := (&universe.Config{}).Read(filesystem.GetFile(universe.FILENAME))
	info_config := (&infocard.Config{}).Read(filesystem.GetFile(infocard.FILENAME, infocard.FILENAME_FALLBACK))
	systems := (&systems.Config{}).Read(universe_config, filesystem)
	market_ships_config := (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_SHIPS))
	market_commodities := (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_COMMODITIES))
	market_misc := (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_MISC))

	// denormalize
	base_denormalizer := (&denormalizer.Denormalizer{}).Create(universe_config)
	base_denormalizer.ReadBaseNames(market_ships_config, universe_config, info_config)
	base_denormalizer.ReadRecycle(market_ships_config, universe_config, systems)
	base_denormalizer.MarketWrite(market_commodities)
	base_denormalizer.MarketWrite(market_ships_config)
	base_denormalizer.MarketWrite(market_misc)
	base_denormalizer.UniverseWrite(universe_config)

	// write
	universe_config.Write().WriteLines(dry_run)
	market_ships_config.Write().WriteLines(dry_run)
	market_commodities.Write().WriteLines(dry_run)
	market_misc.Write().WriteLines(dry_run)
	system_files := systems.Write()
	for _, system_file := range system_files {
		system_file.WriteLines(dry_run)
	}

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
