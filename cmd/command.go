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
		Usage:       "--path=./config/config.toml --config=./build/config",
		Description: "start sky",
		Action:      starts,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "this sky config",
			},
			&cli.StringFlag{
				Name:        "path",
				Usage:"path",
			},
		},
	}
)

func starts(cli *cli.Context) error {
	config := cli.String("config")
	if config == "" {
		return errors.New("please add --config=./config/config.toml")
	}
	path := cli.String("path")
	if err := server.Sky(config,path); err != nil {
		fmt.Println("run sky failed ", err)
		return err
	}
	return nil
}
