package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func RunCli() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "version",
				Usage:   "show version",
				Aliases: []string{"v"},
			},
			&cli.StringFlag{
				Name:    "lang",
				Value:   "english",
				Usage:   "select `language`",
				Aliases: []string{"l"},
				Action: func(ctx *cli.Context, v string) error {
					if v != "english" && v != "indonesia" {
						return fmt.Errorf("invalid language. Please use 'english' or 'indonesia'")
					}
					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			if cCtx.Bool("version") {
				fmt.Println("Version: 0.0.1")
				return nil
			}

			if cCtx.String("lang") == "indonesia" {
				fmt.Println("Halo Dunia!")
			} else {
				fmt.Println("Hello World!")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
