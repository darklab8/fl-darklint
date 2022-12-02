package settings

import "os"

var FreelancerFolderLocation string
var TestingIntegration bool = false

func init() {
	FreelancerFolderLocation = os.Getenv("DARKTOOL_FREELANCER_FOLDER")
	if len(FreelancerFolderLocation) == 0 {
		FreelancerFolderLocation = "/home/naa/repos/pet_projects/darklab_freelancer_darktool/Discovery-DEV-Groshyr"
	}
}
