package utils

import (
	"github.com/sirupsen/logrus"
)

// FailOnError is an error handler that fails the program if an error is passed in
func FailOnError(err error, msg string) {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	if err != nil {
		log.Error("%s: %s", msg, err)
	}
}

// LogOnError is an error handler that logs the error if one is passed in
func LogOnError(err error) {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	if err != nil {
		log.Error("Error: %s", err)
	}
}
