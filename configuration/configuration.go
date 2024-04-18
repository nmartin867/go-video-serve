package config

import (
	"log/syslog"
)

var config *AppConfig

type RuntimeEnvironment int

const (
	Release RuntimeEnvironment = iota
	Development
	Test
)

func (e RuntimeEnvironment) String() string {
	return [...]string{"Release", "Development", "Test"}[e]
}

type AppConfig struct {
	AppName     string
	LogPriority syslog.Priority
	SysLogURI   string
	Port        string
	RuntimeMode RuntimeEnvironment
}

func Init() {
	config = Debug()
}
// TODO: Actually read the config
func Debug() *AppConfig {
	return &AppConfig{
		AppName:     "Video Service",
		LogPriority: syslog.LOG_DEBUG,
		SysLogURI:   "localhost:1514",
		Port:        "8080",
		RuntimeMode: Development,
	}
}

//func parsePriority() syslog.Priority {
//	val := os.Getenv("LOG_PRIORITY")
//	n, err := strconv.Atoi(val)
//	if err != nil {
//		logrus.Warn("Invalid config value:", "LOG_PRIORITY")
//		return syslog.LOG_WARNING
//	}
//	return syslog.Priority(n)
//}
