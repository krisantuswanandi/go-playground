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
		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"g"},
				Usage:   "Say hello to the provided `name`",
				Action: func(ctx *cli.Context) error {
					if ctx.String("lang") == "indonesia" {
						fmt.Println("Halo,", ctx.Args().First()+"!")
					} else {
						fmt.Println("Hello,", ctx.Args().First()+"!")
					}
					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("version") {
				fmt.Println("Version: 0.0.1")
				return nil
			}

			fmt.Println("Welcome to the app!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
