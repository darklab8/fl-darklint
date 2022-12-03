package filefind

import (
	"darktool/settings"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoverFiles(t *testing.T) {
	// Write some data example in order to remove integration flag
	_, filename, _, _ := runtime.Caller(0)
	directory := filepath.Dir(filename)
	test_directory := filepath.Join(directory, "testdata")
	settings.FreelancerFolderLocation = test_directory
	DiscoverConfigs()

	assert.Equal(t, 2, len(Filesystem.Files), "expected 2 files, fount smth else")
}
