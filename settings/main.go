package settings

import (
	"os"
	"strings"

	"darktool/settings/loglevel"
	_ "embed"

	log "github.com/sirupsen/logrus"
)

var FreelancerFreelancerLocation string
var TestingIntegration bool = false
var Debug bool = false
var DryRun = false
var LogLevel = loglevel.Warning

//go:embed version.txt
var Version string

func init() {
	log.Info("init settings")

	// =========== NORMAL SETTINGS ==================
	if path, ok := os.LookupEnv("FLDARKLINT_PROJECT_FOLDER"); ok {
		FreelancerFreelancerLocation = path
	} else {
		exe_path, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		FreelancerFreelancerLocation = exe_path
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

	// =========== ENVIRONMENT OVERRIDES ============
	if len(os.Getenv("DARKTOOL_FREELANCER_FOLDER")) > 0 {
		FreelancerFreelancerLocation = os.Getenv("DARKTOOL_FREELANCER_FOLDER")
	}

	if len(os.Getenv("TEST_INTEGRATION")) != 0 {
		TestingIntegration = true
	}

	// Debug override
	Debug = strings.Compare(os.Getenv("DARKTOOL_DEBUG"), "") != 0
	// Debug logging
	log.SetLevel(log.DebugLevel)
}
