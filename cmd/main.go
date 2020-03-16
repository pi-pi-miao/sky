package main

import (
	"github.com/urfave/cli"
	"sky/pkg/sky"
	"time"
	"errors"
	"fmt"
	"os"
)

var (
	version = "0.1"

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
	reload = cli.Command{
		Name:                   "reload",
		Usage:                  "",
		Description:            "reload sky",
		Action:                 reloads,
		Flags:                  []cli.Flag{
			&cli.StringFlag{
				Name:"",
				Usage:"",
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
	internalAddr := cli.String("internal")
	if addr == "" {
		return errors.New("please add --internal_addr=ip:port")
	}
	if err := sky.Sky(addr,internalAddr);err != nil {
		fmt.Println("run sky failed ",err)
		return err
	}
	return nil
}

func reloads(cli *cli.Context)error{
	return nil
}

func main(){
	aspired := cli.NewApp()
	aspired.Name = "sky"
	aspired.Compiled = time.Now()
	aspired.Version = version
	aspired.Commands = []cli.Command{
		start,
	}
	if err := aspired.Run(os.Args);err != nil {
		fmt.Println("sky start failed",err)
	}
}