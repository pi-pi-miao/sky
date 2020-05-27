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


func InitAll()(err error){
	Sky = &Config{SkyConfig:&SkyConfig{
		Stop:make(chan struct{}),
	}}
	apis.Apis()
	go pprofServer()
	if err := initClient();err != nil {
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


func pprofServer(){
	http.ListenAndServe(Sky.SkyConfig.SkyPProfAddr, nil)
}