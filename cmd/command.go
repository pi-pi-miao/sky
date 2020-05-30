package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"sky/pkg/server"
)

var (
	start = cli.Command{

		Name:        "start",
		Usage:       "--config=./config/config.toml",
		Description: "start sky",
		Action:      starts,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "this sky config",
			},
		},
	}
)

func starts(cli *cli.Context) error {
	config := cli.String("config")
	if config == "" {
		return errors.New("please add --config=./config/config.toml")
	}
	if err := server.Sky(config); err != nil {
		fmt.Println("run sky failed ", err)
		return err
	}
	return nil
}
