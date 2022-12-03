package settings

import (
	"os"
	"strings"

	"darktool/settings/loglevel"

	log "github.com/sirupsen/logrus"
)

var FreelancerFolderLocation string
var TestingIntegration bool = false
var Debug bool = false

var LogLevel = loglevel.Warning

func init() {
	FreelancerFolderLocation = os.Getenv("DARKTOOL_FREELANCER_FOLDER")
	if len(FreelancerFolderLocation) == 0 {
		FreelancerFolderLocation = "/home/naa/repos/pet_projects/darklab_freelancer_darktool/Discovery-DEV-Groshyr"
	}

	if len(os.Getenv("TEST_INTEGRATION")) != 0 {
		TestingIntegration = true
	}

	// Enabling log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	switch LogLevel {
	case loglevel.Panic:
		log.SetLevel(log.PanicLevel)
	case loglevel.Fatal:
		log.SetLevel(log.ErrorLevel)
	case loglevel.Warning:
		log.SetLevel(log.WarnLevel)
	case loglevel.Info:
		log.SetLevel(log.InfoLevel)
	case loglevel.Debug:
		log.SetLevel(log.DebugLevel)
	}

	// Debug override
	Debug = strings.Compare(os.Getenv("DARKTOOL_DEBUG"), "") != 0
	// Debug logging
	log.SetLevel(log.DebugLevel)
}
