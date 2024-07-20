/*
Find duplicates in a file by regular expression
*/
package findduplicates

import (
	"fmt"
	"os"
	"regexp"

	"github.com/darklab8/fl-darklint/darklint/settings/logus"

	"github.com/darklab8/go-utils/utils/utils_logus"
	"github.com/darklab8/go-utils/utils/utils_types"
)

func regexCompile(expression string) *regexp.Regexp {
	var err error

	regex, err := regexp.Compile(expression)
	logus.Log.CheckFatal(err, "failed to compile regex expression "+expression)
	return regex
}

func Main(path utils_types.FilePath, regex string) {
	regx := regexCompile(regex)
	file, err := os.ReadFile(string(path))
	logus.Log.CheckFatal(err, "failed to read file ", utils_logus.FilePath(path))
	content := string(file)

	foundlines := regx.FindAllString(content, -1)

	counter := make(map[string]int, 0)

	for _, line := range foundlines {
		counter[line] = counter[line] + 1
	}

	for key, value := range counter {
		if value < 2 {
			continue
		}
		fmt.Printf("%s (N %d times)\n", key, value)
	}
}
