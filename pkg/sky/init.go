package sky

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
	_ "net/http/pprof"
	"sky/apis"
	"sky/pkg/client"
	logger "sky/pkg/log"
)


func InitAll(path string)(err error){
	Sky = &Config{SkyConfig:&SkyConfig{
		Stop:make(chan struct{}),
	}}
	apis.Apis()
	if err := initConfig(path);err != nil {
		return err
	}
	if err := initClient();err != nil {
		return err
	}
	if err := initLog();err != nil {
		return err
	}
	if err := pprofServer();err != nil {
		return err
	}
	if err := run();err != nil {
		return err
	}
	return nil
}

func initConfig(path string)error{
	if _, err := toml.DecodeFile(path, Sky); err != nil {
		fmt.Printf("toml decode err %v", err)
		return err
	}
	return nil
}

func initLog()error{
	if err := logger.InitLogger(Sky);err != nil {
		return err
	}
	return nil
}

func initClient()(err error){
	Sky.SkyConfig.Client,err = client.GetClient()
	return err
}


func pprofServer()error{
	ch := make(chan error)
	go func(){
		if err := http.ListenAndServe(Sky.SkyConfig.SkyPProfAddr, nil);err != nil {
			ch <- err
			fmt.Println("start pprofServer err ",err)
		}
	}()
	select {
	case err :=<- ch:
		close(ch)
		return err
	default:
		return nil
	}
}

func run()error{
	if err := http.ListenAndServe(Sky.SkyConfig.SkyAddr,nil);err != nil {
		return err
	}
	return nil
}