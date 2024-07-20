package settings

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/darklab8/fl-darklint/darklint/settings/logus"
	"github.com/darklab8/go-utils/utils/enverant"
	"github.com/darklab8/go-utils/utils/utils_settings"
)

//go:embed version.txt
var Version string

var ToolName string = "darklint"
var ToolNameCap string = strings.ToUpper(ToolName)

type DarklingEnvVars struct {
	utils_settings.UtilsEnvs
	configs_settings.ConfEnvVars
}

var Env DarklingEnvVars

func init() {
	env := enverant.NewEnverant()
	Env = DarklingEnvVars{
		UtilsEnvs:   utils_settings.GetEnvs(env),
		ConfEnvVars: configs_settings.GetEnvs(env),
	}
	fmt.Sprintln("conf=", Env)
}

func init() {
	logus.Log.Info("init settings")
}
