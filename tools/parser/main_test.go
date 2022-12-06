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
	Run(dry_run)
}
