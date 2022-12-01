/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
