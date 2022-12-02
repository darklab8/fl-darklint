package parser

import (
	"darktool/settings"
	"testing"
)

func TestSimple(t *testing.T) {
	if !settings.TestingIntegration {
		return
	}
	Run()
}
