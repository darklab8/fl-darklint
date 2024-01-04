package settings

import (
	"os"
	"strings"

	"darklint/fldarklint/settings/logus"
	_ "embed"
)

var FreelancerFreelancerLocation string

var TestingIntegration bool = false

//go:embed version.txt
var Version string

var ToolName string = "fldarklint"
var ToolNameCap string = strings.ToUpper(ToolName)

func init() {
	logus.Log.Info("init settings")

	// =========== NORMAL SETTINGS ==================
	if path, ok := os.LookupEnv(ToolNameCap + "_PROJECT_FOLDER"); ok {
		FreelancerFreelancerLocation = path
	} else {
		exe_path, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		FreelancerFreelancerLocation = exe_path
	}

	if _, ok := os.LookupEnv("TEST_INTEGRATION"); ok {
		TestingIntegration = true
	}

}
