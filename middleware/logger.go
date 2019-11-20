package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
    rotatelogs "github.com/lestrrat/go-file-rotatelogs"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

func LoggerInit(c *gin.Engine, file string) {
    logger, err := loggerMiddleware(file)
    if err != nil {
        panic(err)
    }
    c.Use(logger)
    loggerError, err := loggerErrorMiddleware(file)
    if err != nil {
        panic(err)
    }
    c.Use(loggerError)
}

func loggerMiddleware(file string) (logger gin.HandlerFunc, err error) {
    logClient := logrus.New()

    if _, err = os.Stat(file); err != nil {
        if _, err = os.Create(file); err != nil {
            panic(err)
        }
    }
    src, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR, os.ModeAppend)
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
        logrus.InfoLevel: logWriter,
        logrus.WarnLevel:  logWriter,
        logrus.ErrorLevel: logWriter,
        logrus.FatalLevel: logWriter,
    }

    lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
    logClient.AddHook(lfHook)

    stdOut := lfshook.NewHook(os.Stdout, &logrus.TextFormatter{})
    logClient.AddHook(stdOut)

    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        end := time.Now()
        latency := end.Sub(start)

        path := c.Request.URL.Path
        clientIP := c.ClientIP()
        method := c.Request.Method
        statusCode := c.Writer.Status()
        logClient.WithFields(logrus.Fields{
            "status_coude": statusCode,
            "latency":      latency,
            "client_ip":    clientIP,
            "method":       method,
            "path":         path,
        }).Info()
    }, err
}

func loggerErrorMiddleware(file string) (logger gin.HandlerFunc, err error) {
    logClient := logrus.New()

    if _, err = os.Stat(file); err != nil {
        if _, err = os.Create(file); err != nil {
            panic(err)
        }
    }
    src, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR, os.ModeAppend)
    if err != nil {
        fmt.Println("err:", err)
        return
    }

    logClient.Out = src
    logClient.SetLevel(logrus.DebugLevel)
    apiErrorLogPath := file + "-error"

    logErrorWriter, err := rotatelogs.New(
        apiErrorLogPath+".%Y-%m-%d-%H-%M.error.log",
        rotatelogs.WithLinkName(apiErrorLogPath),
        rotatelogs.WithMaxAge(7*24*time.Hour),
        rotatelogs.WithRotationTime(24*time.Hour),
    )

    writeErrorMap := lfshook.WriterMap{
        logrus.WarnLevel:  logErrorWriter,
        logrus.ErrorLevel: logErrorWriter,
        logrus.FatalLevel: logErrorWriter,
    }

    errorHook := lfshook.NewHook(writeErrorMap, &logrus.JSONFormatter{})
    logClient.AddHook(errorHook)

    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        end := time.Now()
        latency := end.Sub(start)

        path := c.Request.URL.Path
        clientIP := c.ClientIP()
        method := c.Request.Method
        statusCode := c.Writer.Status()
        logClient.WithFields(logrus.Fields{
            "status_coude": statusCode,
            "latency":      latency,
            "client_ip":    clientIP,
            "method":       method,
            "path":         path,
        }).Error()
    }, err
}
