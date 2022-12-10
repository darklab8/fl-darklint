package denormalizer

import (
	"darktool/tools/parser"
)

func Run(parsed *parser.Parsed) {
	(&BaseDenormalizer{}).Read(parsed).Write(parsed)
}
