package randrotator

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

type Input struct {
	Delimiter      string
	RoundedNumbers *int
}

func Run(input Input) string {
	log.Info("randrotator: Run - start")
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

	log.Info("randrotator: Run - finished")
	return fmt.Sprintf("%v\n", sb.String())
}
