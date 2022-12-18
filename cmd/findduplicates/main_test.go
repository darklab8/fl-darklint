package findduplicates

import (
	"darktool/tools/utils"
	"path/filepath"
	"testing"
)

func TestMain(t *testing.T) {
	folder := utils.GetCurrrentTestFolder()
	testfile := filepath.Join(folder, "example.txt")
	Main(testfile, "nickname = .*")
}
