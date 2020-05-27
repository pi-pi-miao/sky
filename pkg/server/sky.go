package server

import (
	"net/http"
	"sky/pkg/initiaze"
)

var (
	ch = make(chan error,2)
)

func Sky(addr,internal string)error{
	initiaze.InitAll()
	if err := http.ListenAndServe(addr,nil);err != nil {
		return err
	}
	return nil
}


