package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"sky/pkg/sky"
)

var (
	start = cli.Command{

		Name:                   "start",
		Usage:					"--addr=0.0.0.0:10001",
		Description:            "start sky",
		Action:                 starts,
		Flags:                  []cli.Flag{
			&cli.StringFlag{
				Name:"addr",
				Usage:"this sky addr",
			},
			&cli.StringFlag{
				Name:"internal",
			},
		},
	}
)

func starts(cli *cli.Context)error{
	addr := cli.String("addr")
	if addr == "" {
		return errors.New("please add --addr=ip:port")
	}
	if err := server.Sky(addr);err != nil {
		fmt.Println("run sky failed ",err)
		return err
	}
	return nil
}
