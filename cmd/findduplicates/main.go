/*
Find duplicates in a file by regular expression
*/
package findduplicates

import (
	"darktool/tools/utils"
	"fmt"
	"io/ioutil"
	"regexp"
)

func regexCompile(expression string) *regexp.Regexp {
	var err error

	regex, err := regexp.Compile(expression)
	utils.CheckFatal(err, "failed to compile regex expression ", expression)
	return regex
}

func Main(path string, regex string) {
	regx := regexCompile(regex)
	file, err := ioutil.ReadFile(path)
	utils.CheckFatal(err, "failed to read file ", path)
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
