package market1ships

import (
	"darktool/tools/parser/file1discovery"
	"darktool/tools/utils"
	"regexp"
	"strconv"
	"strings"
)

type MarketGood struct {
	Name   string
	Values []float32
}

type BaseGood struct {
	Name  string
	Base  string
	Goods []MarketGood
}

var MarketShips struct {
	Intro      []string
	Base_goods []*BaseGood
}

func Parse() {
	file := utils.File.OpenToReadF(utils.File{Filepath: file1discovery.Filesystem.Hashmap["market_ships.ini"].AbsPath})
	defer file.Close()

	lines := file.ReadLines()
	// comments
	comment_exp, err := regexp.Compile(`;(.*)`)
	utils.CheckPanic(err)
	// [BaseGood]
	base_group_ext, err := regexp.Compile(`^\[.*\]`)
	utils.CheckPanic(err)
	// `base = br01_01_base` or `MarketGood = dsy_arrow_package, 1, -1, 1, 1, 0, 1, 1`
	param_exp, err := regexp.Compile(`^([a-zA-Z_]+)\s=\s([a-zA-Z_, 0-9-]+)`)
	utils.CheckPanic(err)

	is_intro := true
	var current_base_good *BaseGood = &BaseGood{}
	for _, line := range lines {

		comment_match := comment_exp.FindStringSubmatch(line)
		base_group_match := base_group_ext.FindStringSubmatch(line)
		param_match := param_exp.FindStringSubmatch(line)

		if len(comment_match) > 0 {
			if is_intro {
				MarketShips.Intro = append(MarketShips.Intro, comment_match[1])
			} else {
				current_base_good.Name = comment_match[1]
			}
			continue
		} else if is_intro {
			is_intro = false
			continue
		}

		if len(base_group_match) > 0 {
			current_base_good = &BaseGood{}
			MarketShips.Base_goods = append(MarketShips.Base_goods, current_base_good)
			continue
		}

		if len(param_match) > 0 {
			if strings.Compare(param_match[1], "base") == 0 {
				current_base_good.Base = param_match[2]
			} else if strings.Compare(param_match[1], "MarketGood") == 0 {
				params := strings.Split(param_match[2], ", ") // data example: dsy_arrow_package, 1, -1, 1, 1, 0, 1, 1
				var floats []float32

				for _, string_number := range params {
					parsed_float, _ := strconv.ParseFloat(string_number, 32)
					floats = append(floats, float32(parsed_float))
				}

				current_base_good.Goods = append(current_base_good.Goods, MarketGood{Name: params[0], Values: floats})
			}
			continue
		}
	}
}
