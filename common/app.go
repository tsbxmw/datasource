package common

import (
	"context"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"time"
)

func App(serviceName string, serviceUsage string, httpServer HttpServer, serviceConfig ServiceConfig) (app *cli.App,
	err error) {
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
					//fmt.Println("Loading config from", config)
					//err, configReal := serviceConfig.ConfigFromFileName(config)
					//if err != nil {
					//	panic(err)
					//}
					httpReal := httpServer.Init(serviceConfig, config)
					go httpReal.Serve()

					// Wait for interrupt signal to gracefully shutdown the server with
					// a timeout of 5 seconds.
					quit := make(chan os.Signal)
					signal.Notify(quit, os.Interrupt, os.Kill)
					<-quit
					httpReal.Shutdown()
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					ctx.Done()
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
					//err, configReal := serviceConfig.ConfigFromFileName(config)
					//if err != nil {
					//	panic(err)
					//}
					httpReal := httpServer.Init(serviceConfig, config)
					go httpReal.ServeWorker()

					// Wait for interrupt signal to gracefully shutdown the server with
					// a timeout of 5 seconds.
					quit := make(chan os.Signal)
					signal.Notify(quit, os.Interrupt)
					<-quit
					//log.Println("Shutdown Server : <<<", conf.ServiceName, ">>>")

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					ctx.Done()
					//log.Println("Server <<<", conf.ServiceName, ">>> Worker Exit  OK")
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
