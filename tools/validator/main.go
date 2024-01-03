/*
Scanned info with package `parser` we try here to validate for being correct
And even suggesting autofixes to Freelancer config files
*/
package validator

import (
	"darktool/settings"
	"darktool/tools/parser"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func Run() {

	_, err := os.Stat(filepath.Join(settings.FreelancerFreelancerLocation, "DATA"))

	fmt.Println(err)
	if os.IsNotExist(err) {
		log.Fatal("freelancer folder is not detected at path=", settings.FreelancerFreelancerLocation, " because DATA folder was not found")
	}

	data := (&parser.Parsed{}).Read(settings.FreelancerFreelancerLocation)

	// see README.go in denormalizer why it was commented out but not removed.
	// denormalizer.Run(data)

	data.Write(settings.DryRun)
}
