package main

import (
    "github.com/sirupsen/logrus"
    "github.com/tsbxmw/datasource/common"
    "github.com/tsbxmw/datasource/data/transport/http"
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
    if err := app.Run(os.Args); err != nil {
        logrus.Error(err)
        panic(err)
    }
}
