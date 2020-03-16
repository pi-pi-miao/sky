package server

import (
	"fmt"
	"net"
	"net/http"
	"sky/util/json"
)

func Login(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case http.MethodPost:
		if err := login(r);err != nil {
			if err != nil {
				Response(w, "10001", fmt.Sprintf("%v", err))
				return
			}
			Response(w, "10000", "login succ")
		}
	default:
		Response(w, "10002", fmt.Sprintf("%v", "request not right"))
	}
	return
}

func Register(w http.ResponseWriter,r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := register(r); err != nil {
			if err != nil {
				Response(w, "10001", fmt.Sprintf("%v", err))
				return
			}
			Response(w, "10000", "succ")
		}
	default:
		Response(w, "10002", fmt.Sprintf("%v", "request not right"))
	}
	return
}

func List(w http.ResponseWriter,r *http.Request){
	switch  {
	case r.Method == http.MethodGet:
		data,err := list()
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		service := ListService{
			ServiceCm:make([]ServiceCm,0,1024),
		}

		for k,_ := range data.Items {
			cm := ServiceCm{
				Data: make(map[string]string,1024),
			}
			cm.Name = data.Items[k].Name
			cm.Data = data.Items[k].Data
			service.ServiceCm = append(service.ServiceCm,cm)
		}
		result,err := json.Marshal(service)
		if err != nil {
			fmt.Println("marshal err",err)
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10001",string(result))
		return
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
		return
	}
}

func GetService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodGet :
		fmt.Println("get service")
		data,err := getService("get",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10000",fmt.Sprintf("%v",data))
		return
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
	}
	return
}

func Response(w http.ResponseWriter,code string,data string){
	result,_ := json.Marshal(Result{
		Code: code,
		Data: data,
	})
	w.Write(result)
}

func CreateService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPost:
		_,err := getService("create",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10000",fmt.Sprintf("%v","succ"))
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
	}
	return
}

func AddService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPost :
		_,err := getService("add_service",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10000",fmt.Sprintf("%v","succ"))
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
	}
	return
}

func AddData(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPost :
		_,err := getService("add",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
			Response(w,"10000",fmt.Sprintf("%v","succ"))
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
	}
	return
}

func UpdateService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodPatch :
		_,err := getService("update",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10000",fmt.Sprintf("%v","succ"))
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
	}
	return
}

func DelService(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodDelete :
		_,err := getService("del_service",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10000",fmt.Sprintf("%v","succ"))
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
	}
	return
}

func DelData(w http.ResponseWriter,r *http.Request){
	switch {
	case r.Method == http.MethodDelete :
		_,err := getService("del_data",r)
		if err != nil {
			Response(w,"10001",fmt.Sprintf("%v",err))
			return
		}
		Response(w,"10000",fmt.Sprintf("%v","succ"))
	default:
		Response(w,"10002",fmt.Sprintf("%v","request not right"))
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
