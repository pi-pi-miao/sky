package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sky/pkg/json"
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

func login(r *http.Request)error{
	s := &User{}
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("read request body err",err)
		return err
	}
	err = json.Unmarshal(body,s)
	if err != nil {
		fmt.Println("unmarshal err",err)
		return err
	}

	// todo check and store
	return nil
}

func register(r *http.Request)error{
	s := &User{}
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("read request body err",err)
		return err
	}
	err = json.Unmarshal(body,s)
	if err != nil {
		fmt.Println("unmarshal err",err)
		return err
	}
	// todo store
	return nil
}