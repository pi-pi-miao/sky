package apis

import (
	"net/http"
	"sky/pkg/server"
)

func Apis(){
	http.HandleFunc("/get", server.GetService)
	http.HandleFunc("/create", server.CreateService)
	http.HandleFunc("/add_service", server.AddService)
	http.HandleFunc("/add_data", server.AddData)
	http.HandleFunc("/updata", server.UpdateService)
	http.HandleFunc("/del_service", server.DelService)
	http.HandleFunc("/del_data", server.DelData)
	http.HandleFunc("login", server.Login)
	http.HandleFunc("register", server.Register)
}