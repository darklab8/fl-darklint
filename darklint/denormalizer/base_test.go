package denormalizer

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darklint/darklint/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := utils_filepath.Dir(utils_filepath.Dir(folder))
	logus.Log.Debug("beginning test", utils_logus.FilePath(freelancer_folder))

	parsed := (&configs_mapped.MappedConfigs{}).Read(freelancer_folder)
	(&BaseDenormalizer{}).Read(parsed).Write(parsed)
	market_ship_lines := parsed.MarketShips.Write().GetLines()

	// isRecycleCandidate
	isRecyclePresent := false
	lines := market_ship_lines
	for _, line := range lines {
		if strings.Contains(line, BASE_KEY_RECYCLE) {
			isRecyclePresent = true
		}
	}

	dry_run := configs_mapped.IsDruRun(true)
	parsed.Write(dry_run)

	assert.True(t, isRecyclePresent)
}
