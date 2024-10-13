package logger

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func GetLogger() *log.Logger {
	return log.StandardLogger()
}
