package common

import (
    "fmt"
    "github.com/gin-gonic/gin"
    rotatelogs "github.com/lestrrat/go-file-rotatelogs"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

func Logger(file string) (logger gin.HandlerFunc, err error) {
    logClient := logrus.New()
    src, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        fmt.Println("err:", err)
        return
    }

    logClient.Out = src
    logClient.SetLevel(logrus.DebugLevel)
    apiLogPath := file

    logWriter, err := rotatelogs.New(
        apiLogPath+".%Y-%m-%d-%H-%M.log",
        rotatelogs.WithLinkName(apiLogPath),
        rotatelogs.WithMaxAge(7*24*time.Hour),
        rotatelogs.WithRotationTime(24*time.Hour),
    )
    writeMap := lfshook.WriterMap{
        logrus.InfoLevel:  logWriter,
        logrus.FatalLevel: logWriter,
    }

    lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
    logClient.AddHook(lfHook)

    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        end := time.Now()
        latency := end.Sub(start)

        path := c.Request.URL.Path
        clientIP := c.ClientIP()
        method := c.Request.Method
        statusCode := c.Writer.Status()
        logClient.Infof("|%3d|%13v|%15s|%s %s|", statusCode, latency, clientIP, method, path)
    }, err
}
