package denormalizer

// import (
// 	"github.com/darklab8/fl-darklint/fldarklint/parser"
// 	"github.com/darklab8/fl-darklint/fldarklint/parser/freelancer/data/equipment/market"
// 	"github.com/darklab8/go-utils/goutils/utils"
// 	"strings"

// 	"path/filepath"
// 	"testing"

// 	log "github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSaveRecycleParams(t *testing.T) {
// 	folder := utils.GetCurrentFolder()
// 	freelancer_folder := filepath.Dir(filepath.Dir(filepath.Dir(folder)))
// 	log.Debug(freelancer_folder)

// 	parsed := (&parser.Parsed{}).Read(freelancer_folder)
// 	(&BaseDenormalizer{}).Read(parsed).Write(parsed)
// 	market_ship_lines := parsed.Market_ships_config.Write().GetLines()

// 	// isRecycleCandidate
// 	isRecyclePresent := false
// 	lines := market_ship_lines
// 	for _, line := range lines {
// 		if strings.Contains(line, market.KEY_RECYCLE) {
// 			isRecyclePresent = true
// 		}
// 	}

// 	dry_run := true
// 	parsed.Write(dry_run)

// 	assert.True(t, isRecyclePresent)
// }
