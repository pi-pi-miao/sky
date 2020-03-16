package sky

import (
	"fmt"
	"net/http"
	"sky/apis"
	"sky/pkg/initiaze"
	"sky/pkg/server"
)

var (
	ch = make(chan error,2)
)

// todo  shut add pprof
func Sky(addr,internal string)error{
	apis.Apis()
	initiaze.InitAll()
	go startHttp(addr)
	select {
	case err :=<- ch:
		return err
	default:
		if err := server.Start(internal);err != nil {
			return err
		}
	}
	return nil
}

func startHttp(addr string){
	defer func() {
		if err := recover();err != nil {
			fmt.Println("[startHttp] goroutine panic")
		}
	}()
	if err := http.ListenAndServe(addr,nil);err != nil {
		ch <- err
		return
	}
}

