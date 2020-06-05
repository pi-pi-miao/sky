package sky

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"k8s.io/client-go/tools/cache"
	"net/http"
	_ "net/http/pprof"
	"sky/pkg/client"
	"sky/pkg/logger"
)

const (
	Namespace = "Namespace"
	Version   = "v1"
)

func InitAll(config,path string) (err error) {
	if err = initSky(config);err != nil {
		return err
	}
	if err = initConfig(path); err != nil {
		return
	}
	if err = initLog(); err != nil {
		return
	}
	if err = pprofServer(); err != nil {
		return
	}
	if err = Sky.SkyConfig.CreateInformer();err != nil {
		return
	}
	if err = Sky.SkyConfig.CheckNs();err != nil {
		return
	}
	if err = Sky.SkyConfig.CheckUser();err != nil {
		return
	}
	if err = run(); err != nil {
		return err
	}
	return nil
}

func initSky(config string) (err error){
	Sky = &Config{
		SkyConfig: &SkyConfig{
			Stop:      make(chan struct{}),
			Informers: make([]cache.InformerSynced, 0,10),
		},
		UserLabels: make(map[string]UserLabels, 10),
		User:       &User{},
		Code:&Code{},
	}
	Sky.SkyConfig.Config = Sky
	if err = initClient(config); err != nil {
		return
	}
	return
}

func initConfig(path string) error {
	if _, err := toml.DecodeFile(path, Sky); err != nil {
		fmt.Printf("toml decode err %v", err)
		return err
	}
	return nil
}

func initLog() error {
	if err := logger.InitLogger(Sky); err != nil {
		return err
	}
	return nil
}

func initClient(path string) (err error) {
	Sky.SkyConfig.Client, err = client.GetClient(path)
	return err
}

func pprofServer() error {
	ch := make(chan error)
	go func() {
		if err := http.ListenAndServe(Sky.SkyConfig.SkyPProfAddr, nil); err != nil {
			ch <- err
			fmt.Println("start pprofServer err ", err)
		}
	}()
	select {
	case err := <-ch:
		close(ch)
		return err
	default:
		return nil
	}
}

func run() error {
	if err := http.ListenAndServe(Sky.SkyConfig.SkyAddr, nil); err != nil {
		return err
	}
	return nil
}
