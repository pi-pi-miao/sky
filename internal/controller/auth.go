package controller

import (
	"fmt"

	"errors"

	"github.com/asaskevich/govalidator"

	"sky/internal/store"
	"sky/pkg/logger"

	"io/ioutil"
	"net/http"
	"sky/pkg/json"
	"sky/pkg/sky"
)

func Login(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case http.MethodPost:
		if err := login(r);err != nil {
			if err != nil {
				Response(w, "10001", fmt.Sprintf("%v", err))
				return
			}
			Response(w, "10000", "login success")
		}
	default:
		Response(w, "10002", "request not right")
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
			Response(w, "10000", "register success")
		}
	default:
		Response(w, "10002", "request not right")
	}
	return
}

func login(r *http.Request)error{
	person := &User{}
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		logger.Errorf(sky.Sky,"read request body err %v",err)
		return err
	}
	err = json.Unmarshal(body,person)
	if err != nil {
		logger.Errorf(sky.Sky,"unmarshal err %v",err)
		return err
	}
	if res,err := govalidator.ValidateStruct(person);err != nil && !res {
		logger.Errorf(sky.Sky,"request param err %v",err)
		return err
	}
	user,err := store.GetUser(person.Email)
	if err != nil {
		logger.Errorf(sky.Sky,"Illegal user request err:%v",err)
		return err
	}
	storeUser := &User{}
	json.Unmarshal(user,storeUser)
	if storeUser.PassWd != person.PassWd {
		logger.Errorf(sky.Sky,"%v wrong password",person.Email)
		return errors.New(fmt.Sprintf("%v wrong password",person.Email))
	}
	return nil
}

func register(r *http.Request)error{
	person := &User{}
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("read request body err",err)
		return err
	}
	err = json.Unmarshal(body,person)
	if err != nil {
		fmt.Println("unmarshal err",err)
		return err
	}
	if res,err := govalidator.ValidateStruct(person);err != nil && !res {
		logger.Errorf(sky.Sky,"request param err %v",err)
		return err
	}
	data,_ := json.Marshal(person)
	if err := store.CreateUser(person.Email,data);err != nil{
		logger.Errorf(sky.Sky,"save user err %v",err)
		return err
	}
	return nil
}