/*
# Error handling functions
*/
package utils

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

func CheckFatal(err error, msg ...string) {
	if err != nil {
		log.Fatalf(strings.Join(msg, ""))
	}
}
