package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "swaggonmia"
	app.Usage = "Insomnia to Swagger converter"
	app.Version = "1.1.0"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "Nick Wallace",
			Email: "nwallace@fyberstudios.com",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "config, c",
					Usage:    "Load configuration from `FILE`",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "insomnia, i",
					Usage:    "Insomnia JSON `FILE`",
					Required: true,
				},
				&cli.StringFlag{
					Name:  "output, o",
					Value: "yaml",
					Usage: "Output json|yaml",
				},
			},
			Usage: "Generate Swagger documentation",
			Action: func(c *cli.Context) error {
				var insomniaFile = c.String("insomnia")
				var configFile = c.String("config")
				var outputFormat = c.String("output")

				if insomniaFile == "" || configFile == "" {
					cli.ShowCompletions(c)
				}

				if outputFormat == "" {
					outputFormat = "json"
				}

				swagger := &Swagger{}
				swagger.Generate(insomniaFile, configFile, outputFormat)

				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Wrong command %q !", command)
	}
	app.Run(os.Args)
}
