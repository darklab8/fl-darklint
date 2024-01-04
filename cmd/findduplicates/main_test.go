package findduplicates

import (
	"testing"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

func TestMain(t *testing.T) {
	folder := utils.GetCurrrentTestFolder()
	testfile := utils_filepath.Join(folder, "example.txt")
	Main(testfile, "nickname = .*")
}
