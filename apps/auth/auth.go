package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tsbxmw/datasource/auth/transport/http"
	"github.com/tsbxmw/datasource/common"
	"os"
)

func main() {
	httpServer := http.HttpServer{}
	config := common.ServiceConfigImpl{}
	app, err := common.App("auth_v1", "auth service for datasource", httpServer, config)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
		panic(err)
	}
}
