package main

import (
	"datasource/auth/transport/http"
	"datasource/common"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	app, err := common.App("auth", "usage for auth", http.HttpServer{})
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
