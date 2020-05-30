package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sky/version"
	"time"
)

func main() {
	sky := cli.NewApp()
	sky.Name = "sky"
	sky.Compiled = time.Now()
	sky.Version = version.String()
	sky.Commands = []cli.Command{
		start,
	}
	if err := sky.Run(os.Args); err != nil {
		fmt.Println("sky start failed", err)
	}
}
