/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"darktool/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	cmd.Execute()
}
