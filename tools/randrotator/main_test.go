package randrotator

import (
	"fmt"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	Delimiter = ", "
	RoundedNumbers = "2"
	result := Run()
	fmt.Println(result)

	if strings.Count(result, ", ") != 2 {
		t.Error("Delimiter is not discovered in result=", result)
	}
}
