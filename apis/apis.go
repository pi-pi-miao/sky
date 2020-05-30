package apis

import (
	"net/http"
	"sky/internal/controller"
)

func Apis() {
	//http.HandleFunc("/list", controller.List)
	//http.HandleFunc("/get", controller.GetService)
	//http.HandleFunc("/create", controller.CreateService)
	//http.HandleFunc("/add_data", controller.AddData)
	//http.HandleFunc("/update", controller.UpdateService)
	//http.HandleFunc("/del_service", controller.DelService)
	//http.HandleFunc("/del_data", controller.DelData)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
}
