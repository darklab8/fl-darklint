package settings

import (
	"os"
	"strings"
)

var FreelancerFolderLocation string
var TestingIntegration bool = false
var Debug bool = false
var LogEnabled = false

func init() {
	FreelancerFolderLocation = os.Getenv("DARKTOOL_FREELANCER_FOLDER")
	if len(FreelancerFolderLocation) == 0 {
		FreelancerFolderLocation = "/home/naa/repos/pet_projects/darklab_freelancer_darktool/Discovery-DEV-Groshyr"
	}

	if len(os.Getenv("TEST_INTEGRATION")) != 0 {
		TestingIntegration = true
	}

	Debug = strings.Compare(os.Getenv("DARKTOOL_DEBUG"), "") != 0

	LogEnabled = Debug
}
