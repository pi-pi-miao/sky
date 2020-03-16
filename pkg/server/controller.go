package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"sky/util/client"
	"sky/util/json"
	"sort"
	"strconv"
)

var (
	ns = "sky"
	name = "login"
)

func getService(method string,r *http.Request)(result interface{},err error){
	s := &Service{}
	body,err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("read request body err",err)
		return nil,err
	}
	err = json.Unmarshal(body,s)
	if err != nil {
		fmt.Println("unmarshal err",err)
		return nil,err
	}
	cli,err := client.GetClient()
	if err != nil {
		fmt.Println("client is not right")
		return nil,err
	}
	switch  {

	case method == "get":
		data,err := cli.CoreV1().ConfigMaps(ns).Get(s.Name,metav1.GetOptions{})
		if err != nil {
			fmt.Println("cm is not right",err)
			return nil,err
		}
		result = data.Data
		return result,nil
	case method == "create":
		_,err = cli.CoreV1().ConfigMaps(ns).Create(&v1.ConfigMap{
			ObjectMeta:   metav1.ObjectMeta{
				Name:s.Name,
			},
		})
		if err != nil {
			return nil,err
		}
		return nil,nil
	case method == "add":
		data,err := cli.CoreV1().ConfigMaps(ns).Get(s.Name,metav1.GetOptions{})
		if err != nil {
			fmt.Println("cm is not right",err)
			return nil,err
		}
		data.Data[s.Key] = s.Data
		_,err = cli.CoreV1().ConfigMaps(ns).Update(&v1.ConfigMap{
			Data:data.Data,
		})
		if err != nil {
			return nil,err
		}
	case method == "add_service":
		_,err = cli.CoreV1().ConfigMaps(ns).Update(&v1.ConfigMap{
			ObjectMeta:   metav1.ObjectMeta{
				Name:s.Data,
			},
		})
	case method == "update":
		data,err := cli.CoreV1().ConfigMaps(ns).Get(s.Name,metav1.GetOptions{})
		if err != nil {
			fmt.Println("cm is not right",err)
			return nil,err
		}
		if _,ok := data.Data[s.Key];!ok {
			return nil,errors.New(fmt.Sprintf("it is not key:%v and data%v",s.Key,s.Data))
		}
		data.Data[s.Key] = s.Data
		_,err = cli.CoreV1().ConfigMaps(ns).Update(&v1.ConfigMap{
			Data:data.Data,
		})
		if err != nil {
			return nil,err
		}
	case method == "del_service":
		err = cli.CoreV1().ConfigMaps(ns).Delete(s.Name,&metav1.DeleteOptions{})
		if err != nil {
			return nil,err
		}
	case method == "del_data":
		data,err := cli.CoreV1().ConfigMaps(ns).Get(s.Name,metav1.GetOptions{})
		if err != nil {
			fmt.Println("cm is not right",err)
			return nil,err
		}
		if _,ok := data.Data[s.Key];!ok {
			return nil,errors.New(fmt.Sprintf("it is not key:%v and data%v",s.Key,s.Data))
		}
		delete(data.Data,s.Key)
		_,err = cli.CoreV1().ConfigMaps(ns).Update(&v1.ConfigMap{
			Data:data.Data,
		})
		if err != nil {
			return nil,err
		}
	}
	return
}


// cm key is name
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
	cli,err := client.GetClient()
	if err != nil {
		fmt.Println("client is not right")
		return err
	}
	data,err := cli.CoreV1().ConfigMaps(ns).Get(name,metav1.GetOptions{})
	if err != nil {
		fmt.Println("cm is not right",err)
		return err
	}
	if b,err := getUser(data.Data[name],s);b{
		return nil
	}else {
		return err
	}
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
	cli,err := client.GetClient()
	if err != nil {
		fmt.Println("client is not right")
		return err
	}
	data,err := cli.CoreV1().ConfigMaps(ns).Get(name,metav1.GetOptions{})
	if err != nil {
		fmt.Println("cm is not right",err)
		return err
	}
	users := &Users{
		users:make([]*User,0,1024),
	}
	err = json.Unmarshal([]byte(data.Data[name]),users.users)
	if err != nil {
		fmt.Println("[register] unmarshal err",err)
		return err
	}
	users.users = append(users.users,s)
	d,err := json.Marshal(users)
	if err != nil{
		fmt.Println("[register] marshal err",err)
		return err
	}
	data.Data[name] = string(d)
	_,err = cli.CoreV1().ConfigMaps(ns).Update(&v1.ConfigMap{
		Data:data.Data,
	})
	if err != nil {
		fmt.Println("[register] update err",err)
		return err
	}
	return nil
}

func getUser(data string,s *User)(bool,error){
	users := &Users{
		users:make([]*User,0,1024),
	}
	err := json.Unmarshal([]byte(data),users.users)
	if err != nil {
		fmt.Println("[getUser] unmarshal err",err)
		return false,err
	}
	sort.Sort(users)
	l := users.Len()
	userLen, err := strconv.ParseUint(strconv.Itoa(l) ,10, 64)
	if err != nil {
		fmt.Println("[getUser] strconv parseUint err",err)
		return false,err
	}
	sub := binaryFind(users.users,0,userLen,s.Id)
	if sub != -1{
		return users.users[sub].PassWd == s.PassWd,nil
	}
	return false,errors.New("no user")
}

func (u *Users) Len() int           { return len(u.users) }
func (u *Users) Swap(i, j int)      { u.users[i], u.users[j] = u.users[j], u.users[i] }
func (u *Users) Less(i, j int) bool { return u.users[i].Id < u.users[j].Id }

func binaryFind(arr []*User, leftIndex uint64, rightIndex uint64, findVal uint64)int{
	if leftIndex > rightIndex {
		return -1
	}
	middle := (leftIndex + rightIndex) / 2

	if arr[middle].Id > findVal {
		binaryFind(arr, leftIndex, middle - 1, findVal)
	} else if arr[middle].Id < findVal {
		binaryFind(arr, middle + 1, rightIndex, findVal)
	}
	return int(middle)
}

func (m *sky)manager(){
	go func() {
		defer func() {
			if err := recover();err != nil{
				fmt.Println("[manager] is panic",err)
			}
		}()
		for {
			select {
			case <- m.stop:
				return
			default:
			}
			// todo
		}
	}()
}

func (m *sky)Heartbeat(){

}

func watch(){
	cli,err := client.GetClient()
	if err != nil {
		fmt.Println("client is not right")
		return
	}
	for {
		w,err := cli.CoreV1().ConfigMaps(ns).Watch(metav1.ListOptions{
			TypeMeta:            metav1.TypeMeta{},
			LabelSelector:       "",
			FieldSelector:       "",
			Watch:               false,
			AllowWatchBookmarks: false,
			ResourceVersion:     "",
			TimeoutSeconds:      nil,
			Limit:               0,
			Continue:            "",
		})
		if err != nil {
			fmt.Println("watch err",err)
			return
		}
		// todo
		fmt.Println("watch event is ",<- w.ResultChan())
	}
}