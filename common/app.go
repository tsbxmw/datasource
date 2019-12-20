package common

import (
	"github.com/urfave/cli"
)

func App(serviceName string, serviceUsage string, httpServer HttpServer) (app *cli.App, err error) {
	var config string
	app = &cli.App{
		Name:  serviceName,
		Usage: serviceUsage,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config, c",
				Usage:       "Load config from `FILE`",
				Destination: &config,
			},
		},
		Commands: []*cli.Command{
			{
				Name:         "httpserver",
				Aliases:      nil,
				Usage:        "server over http",
				UsageText:    "",
				Description:  "",
				ArgsUsage:    "",
				Category:     "",
				BashComplete: nil,
				Before:       nil,
				After:        nil,
				Action: func(c *cli.Context) error {
					conf, err := ConfigFromFileName(config)
					if err != nil {
						panic(err)
					}
					//httpServer = transport.HttpServerImpl{
					//    SvcName:    conf.ServiceName,
					//    Address:    conf.HttpAddr,
					//    Port:       conf.Port,
					//    DbUri:      conf.DbUri,
					//    ConsulAddr: conf.ConsulAddr,
					//    JaegerAddr: conf.JaegerAddr,
					//}
					httpReal := httpServer.Init(&conf)
					httpReal.Serve()

					return nil
				},
				OnUsageError:       nil,
				Subcommands:        nil,
				Flags:              nil,
				SkipFlagParsing:    false,
				HideHelp:           false,
				Hidden:             false,
				HelpName:           "",
				CustomHelpTemplate: "",
			},
			{
				Name:         "worker-server",
				Aliases:      nil,
				Usage:        "server of worker",
				UsageText:    "",
				Description:  "",
				ArgsUsage:    "",
				Category:     "",
				BashComplete: nil,
				Before:       nil,
				After:        nil,
				Action: func(c *cli.Context) error {
					conf, err := ConfigFromFileName(config)
					if err != nil {
						panic(err)
					}

					httpReal := httpServer.Init(&conf)
					httpReal.ServeWorker()

					return nil
				},
				OnUsageError:       nil,
				Subcommands:        nil,
				Flags:              nil,
				SkipFlagParsing:    false,
				HideHelp:           false,
				Hidden:             false,
				HelpName:           "",
				CustomHelpTemplate: "",
			},
			{
				Name:         "swagger-api",
				Aliases:      nil,
				Usage:        "init swagger apis",
				UsageText:    "",
				Description:  "",
				ArgsUsage:    "",
				Category:     "",
				BashComplete: nil,
				Before:       nil,
				After:        nil,
				Action: func(c *cli.Context) error {
					print("Start swagger api create or init")
					return nil
				},
				OnUsageError:       nil,
				Subcommands:        nil,
				Flags:              nil,
				SkipFlagParsing:    false,
				HideHelp:           false,
				Hidden:             false,
				HelpName:           "",
				CustomHelpTemplate: "",
			},
		},
	}
	return

}
