/*
Copyright © 2022 dd84ai <dd84ai@gmail.com>

A set of tools for gaming community Freelancer Discovery
in order to be more productive during its development configuration
*/
package main

import (
	"darktool/cmd"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	DEBUG := os.Getenv("DEBUG")
	logEnabled := (strings.Compare(DEBUG, "") != 0)
	if logEnabled {

		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.SetLevel(log.DebugLevel)
		log.Info("log enabled")
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	cmd.Execute()
}
