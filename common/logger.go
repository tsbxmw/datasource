package common

import (
    "github.com/sirupsen/logrus"
)

var LogrusLogger *logrus.Logger

func InitLogger() *logrus.Logger {
    if LogrusLogger != nil {
        return LogrusLogger
    }
    LogrusLogger = logrus.New()
    return LogrusLogger
}
