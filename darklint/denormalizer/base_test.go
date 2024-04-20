package denormalizer

import (
	"fmt"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darklint/darklint/settings"
	"github.com/darklab8/fl-darklint/darklint/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"

	"testing"
)

func TestSaveRecycleParams(t *testing.T) {
	freelancer_folder := utils_types.FilePath(settings.FreelancerFreelancerLocation)
	logus.Log.Info("beginning test", utils_logus.FilePath(freelancer_folder))
	fmt.Println(freelancer_folder)

	parsed := configs_mapped.NewMappedConfigs().Read(freelancer_folder)
	NewBaseDenormalizer().Read(parsed).Write(parsed)
	market_ship_lines := parsed.Market.Write()
	_ = market_ship_lines // geting file lines example

	parsed.Write(configs_mapped.IsDruRun(true))
}
