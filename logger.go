package main

import (
	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"os"
)

type RuntimeEnvironment int

const (
	Release RuntimeEnvironment = iota
	Development
	Test
)

func (e RuntimeEnvironment) String() string {
	return [...]string{"Release", "Development", "Test"}[e]
}

func CreateLogger() *logrus.Logger {
	appTag := os.Getenv("APP_NAME")
	syslogUri := os.Getenv("SYSLOG_URL")
	switch env := os.Getenv("development"); env {

	}

	logger := logrus.New()
	hook, err := logrus_syslog.NewSyslogHook("udp", syslogUri, logLevel, appTag)
	return logger
}
