package common


import (
    "github.com/urfave/cli"
)


func App(serviceConfig ServiceConfig) (app *cli.App, err error) {

    app = cli.NewApp()

    app.Name = serviceConfig.ServiceName
    app.Usage = serviceConfig.ServiceName

    var config string

    return
}