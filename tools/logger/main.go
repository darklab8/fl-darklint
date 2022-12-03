package logger

import (
	"darktool/settings"

	log "github.com/sirupsen/logrus"
)

func init() {
	if settings.LogEnabled {
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
		log.SetLevel(log.DebugLevel)
		log.Info("log enabled")
	} else {
		log.SetLevel(log.ErrorLevel)
	}
}
