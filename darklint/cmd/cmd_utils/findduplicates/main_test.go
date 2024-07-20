package findduplicates

import (
	"testing"

	"github.com/darklab8/go-utils/utils/utils_filepath"
	"github.com/darklab8/go-utils/utils/utils_os"
)

func TestMain(t *testing.T) {
	folder := utils_os.GetCurrrentTestFolder()
	testfile := utils_filepath.Join(folder, "example.txt")
	Main(testfile, "nickname = .*")
}
