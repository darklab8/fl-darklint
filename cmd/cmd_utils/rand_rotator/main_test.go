package rand_rotator

import (
	"fmt"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	roundedN := 2
	input := Input{
		Delimiter:      ", ",
		RoundedNumbers: &roundedN,
	}
	result := Run(input)
	fmt.Println(result)

	if strings.Count(result, ", ") != 2 {
		t.Error("Delimiter is not discovered in result=", result)
	}
}
