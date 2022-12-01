package randrotator

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

type InputData struct {
	Delimiter      string
	RoundedNumbers string
}

func Run(input InputData) string {
	rand.Seed(time.Now().UnixNano())

	x := rand.Float64() * 180
	y := rand.Float64() * 180
	z := rand.Float64() * 180
	var sb strings.Builder
	roundedNeeded, _ := strconv.Atoi(input.RoundedNumbers)
	sb.WriteString(fmt.Sprintf("%v", toFixed(x, roundedNeeded)))
	sb.WriteString(fmt.Sprintf("%v", input.Delimiter))
	sb.WriteString(fmt.Sprintf("%v", toFixed(y, roundedNeeded)))
	sb.WriteString(fmt.Sprintf("%v", input.Delimiter))
	sb.WriteString(fmt.Sprintf("%v", toFixed(z, roundedNeeded)))

	return fmt.Sprintf("%v\n", sb.String())
}
