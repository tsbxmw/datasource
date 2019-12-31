package main

import (
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/data/transport/http"
    "fmt"
    "github.com/sirupsen/logrus"
    "os"
)

func main() {
	httpServer := http.HttpServer{}
	config := common.ServiceConfigImpl{}
    app, err := common.App("datasource", "usage for datasource", httpServer, config)
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
