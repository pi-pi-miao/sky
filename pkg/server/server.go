package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func Login(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case http.MethodPost:
		if err := login(r);err != nil {
			io.WriteString(w,fmt.Sprintf("%v",fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v",err),
			})))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",fmt.Sprintf("%v", Result{
			Code: "10001",
			Data: fmt.Sprintf("%v","login succ"),
		})))
	}
}

func Register(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case http.MethodPost:
		if err := register(r);err != nil {
			io.WriteString(w,fmt.Sprintf("%v",fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v",err),
			})))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",fmt.Sprintf("%v", Result{
			Code: "10001",
			Data: fmt.Sprintf("%v","register succ"),
		})))
	}
}

func GetService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodGet :
		data,err := getService("get",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",data))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func CreateService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPost:
		_,err := getService("create",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v","succ"),
			})))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func AddService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPost :
		_,err := getService("add_service",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v","succ"),
			})))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func AddData(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPost :
		_,err := getService("add",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v","succ"),
			})))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func UpdateService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPatch :
		_,err := getService("update",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v","succ"),
			})))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func DelService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodDelete :
		_,err := getService("del_service",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v","succ"),
			})))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func DelData(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodDelete :
		_,err := getService("del_data",r)
		if err != nil {
			io.WriteString(w,fmt.Sprintf("%v", Result{
				Code: "10001",
				Data: fmt.Sprintf("%v",err),
			}))
			return
		}
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10000",
				Data: fmt.Sprintf("%v","succ"),
			})))
	default:
		io.WriteString(w,fmt.Sprintf("%v",
			fmt.Sprintf("%v", Result{
				Code: "10002",
				Data: fmt.Sprintf("%v","request not right"),
			})))
	}
	return
}

func Start(addr string)error{
	l,err := net.Listen("tcp",addr)
	if err != nil {
		fmt.Println("[start] listen err",err)
		return err
	}
	manager := &skyManager{manager: make([]*sky,0,1024)}
	for {
		conn,err := l.Accept()
		if err != nil {
			fmt.Println("[start] accept err ",err)
			return err
		}
		s := &sky{conn: conn,stop:make(chan struct{})}
		manager.manager = append(manager.manager,s)
		s.manager()
	}
}
