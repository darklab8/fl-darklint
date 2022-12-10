package parser

import (
	"darktool/settings"
	"darktool/tools/utils"
	"path/filepath"
	"testing"
)

func TestSimple(t *testing.T) {
	current_folder := utils.GetCurrentFolder()
	settings.FreelancerFreelancerLocation = filepath.Dir(filepath.Dir(current_folder))
	dry_run := true

	parsed := (&Parsed{}).Read(settings.FreelancerFreelancerLocation)
	parsed.Write(dry_run)
}
