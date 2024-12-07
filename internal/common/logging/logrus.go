package logging

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)


func Init() {
	SetFormatter(logrus.StandardLogger())
	logrus.SetLevel(logrus.DebugLevel)
}


func SetFormatter(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",  // 严重
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocal, _ := strconv.ParseBool(os.Getenv("IS_LOCAL")); isLocal {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
}