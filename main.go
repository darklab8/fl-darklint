/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/darklab8/darklab_freelancer_darktool/darktool/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	cmd.Execute()
}
