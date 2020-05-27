package controller

import (
	"net/http"
	"sky/pkg/json"
)

type Service struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Key  string `json:"key"`
}


type Result struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

type ListService struct {
	ServiceCm []ServiceCm  `json:"items"`
}

type ServiceCm struct {
	Name string              `json:"name"`
	Data map[string]string   `json:"data"`
}

//        ----  login

type Users struct {
	users []*User   `json:"users"`
}

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Role Role   `json:"role"`
	Lessee []*Lessee `json:"lessee"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	PassWd string `json:"passwd"`
}

type Role struct {
	Admin string `json:"admin"`
	Audit string `json:"audit"`
}

type Lessee struct {
	Name string  `json:"name"`
}


func Response(w http.ResponseWriter,code string,data string){
	result,_ := json.Marshal(Result{
		Code: code,
		Data: data,
	})
	w.Write(result)
}