/*
Package generating random rotation for object
*/
package rand_rotator

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/darklab8/fl-darklint/darklint/settings/logus"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

/*
input to use this module
For easy binding to Cobra/Viper integration, we define input purely as reference objects
*/
type Input struct {
	// delimiter separating x,y,z
	Delimiter string

	// precision of rounding floats. 2 = 0.34
	RoundedNumbers *int
}

func Run(input Input) string {
	logus.Log.Info("randrotator: Run - start")
	rand.Seed(time.Now().UnixNano())

	x := rand.Float64() * 180
	y := rand.Float64() * 180
	z := rand.Float64() * 180
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%v", toFixed(x, *input.RoundedNumbers)))
	sb.WriteString(fmt.Sprintf("%v", input.Delimiter))
	sb.WriteString(fmt.Sprintf("%v", toFixed(y, *input.RoundedNumbers)))
	sb.WriteString(fmt.Sprintf("%v", input.Delimiter))
	sb.WriteString(fmt.Sprintf("%v", toFixed(z, *input.RoundedNumbers)))

	logus.Log.Info("randrotator: Run - finished")
	return fmt.Sprintf("%v\n", sb.String())
}
