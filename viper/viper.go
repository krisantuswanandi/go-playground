package viper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func Run() {
	viper.SetConfigName(".apprc")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME")
	viper.ReadInConfig()

	app := &cli.App{
		Name:  "app",
		Usage: "read and write to rc file",
		Commands: []*cli.Command{
			{
				Name: "get",
				Action: func(c *cli.Context) error {
					key := c.Args().First()

					if viper.IsSet(key) {
						val := viper.GetString(key)
						fmt.Println("get " + key + ": " + val)
					} else {
						fmt.Println("no variable " + key)
					}
					return nil
				},
			},
			{
				Name: "set",
				Action: func(c *cli.Context) error {
					key := c.Args().First()
					val := c.Args().Get(1)
					fmt.Println("set " + key + " to " + val)
					viper.Set(key, val)
					viper.WriteConfig()
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
