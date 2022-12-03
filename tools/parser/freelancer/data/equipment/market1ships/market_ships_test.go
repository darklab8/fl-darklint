package market1ships

import (
	"darktool/tools/parser/parserutils/filefind"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	filefind.Load()
	Load()
	assert.Greater(t, len(LoadedMarketShips.Base_goods), 0, "market ships sections were not scanned")
	assert.Greater(t, len(LoadedMarketShips.Base_goods[0].Goods), 0, "market ships sections were not scanned")
}
