package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
	"os"
)

func RequestLogger() gin.HandlerFunc {
	appTag := os.Getenv("APP_NAME")
	syslogUri := os.Getenv("SYSLOG_URL")

	log := logrus.New()
	hook, err := logrus_syslog.NewSyslogHook("udp", syslogUri, syslog.LOG_WARNING, appTag)
	if err == nil {
		log.Hooks.Add(hook)
	}
	return func(context *gin.Context) {

	}
}
