package main

import (
    "datasource/common"
    "datasource/transport/http"
    "fmt"
    "github.com/sirupsen/logrus"
    "os"
)

func main() {
    app, err := common.App("datasource", "usage for datasource", http.HttpServer{})
    if err != nil {
        logrus.Error(err)
        panic(err)
    }
    fmt.Println("start main")
    if err := app.Run(os.Args); err != nil {
        logrus.Error(err)
        panic(err)
    }
}
