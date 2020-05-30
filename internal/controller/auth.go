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

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := login(r); err != nil {
				Response(w, sky.Sky.Code.RequestError, fmt.Sprintf("%v", err))
				return
		}else {
			Response(w, sky.Sky.Code.RequestSuccess, "login success")
		}
	default:
		Response(w, sky.Sky.Code.IllegalRequest, "request not right")
	}
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := register(r); err != nil {
				Response(w, sky.Sky.Code.RequestError, fmt.Sprintf("%v", err))
				return
		}else {
			Response(w, sky.Sky.Code.RequestSuccess, "register success")
		}
	default:
		Response(w, sky.Sky.Code.IllegalRequest, "request not right")
	}
	return
}

func login(r *http.Request) error {
	person := &User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf(sky.Sky, "[login] read request body err %v", err)
		return err
	}
	err = json.Unmarshal(body, person)
	if err != nil {
		logger.Errorf(sky.Sky, "[login] unmarshal err %v", err)
		return err
	}
	if res, err := govalidator.ValidateStruct(person); err != nil && !res {
		logger.Errorf(sky.Sky, "[login] request param err %v", err)
		return err
	}
	user, err := store.GetUser(person.Email)
	if err != nil {
		logger.Errorf(sky.Sky, "[login] Illegal user request err:%v", err)
		return err
	}
	storeUser := &User{}
	json.Unmarshal(user, storeUser)
	if storeUser.PassWd != person.PassWd {
		logger.Errorf(sky.Sky, "[login] %v login wrong password", person.Email)
		return errors.New(fmt.Sprintf("%v login wrong password", person.Email))
	}
	return nil
}

func register(r *http.Request) error {
	person := &User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf(sky.Sky, "[register] read request body err %v", err)
		return err
	}
	err = json.Unmarshal(body, person)
	if err != nil {
		logger.Errorf(sky.Sky, "[register] unmarshal err %v", err)
		return err
	}
	if res, err := govalidator.ValidateStruct(person); err != nil && !res {
		logger.Errorf(sky.Sky, "[register] request param err %v", err)
		return err
	}
	data, _ := json.Marshal(person)
	if err := store.CreateUser(person.Email, data); err != nil {
		logger.Errorf(sky.Sky, "[register] save user err %v", err)
		return err
	}
	return nil
}
