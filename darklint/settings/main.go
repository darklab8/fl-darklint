package settings

import (
	"os"
	"strings"

	_ "embed"

	"github.com/darklab8/fl-darklint/darklint/settings/logus"
)

var FreelancerFreelancerLocation string

var TestingIntegration bool = false

//go:embed version.txt
var Version string

var ToolName string = "darklint"
var ToolNameCap string = strings.ToUpper(ToolName)

func init() {
	logus.Log.Info("init settings")

	// =========== NORMAL SETTINGS ==================
	if path, ok := os.LookupEnv(ToolNameCap + "_FREELANCER_FOLDER"); ok {
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
