package initiaze

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sky/util/client"
)

var (
	ns = "sky"
)

func initNs(){
	cli,err := client.GetClient()
	if err != nil {
		fmt.Println("get cli err",err)
		return
	}
	_,err = cli.CoreV1().Namespaces().Get(ns,metav1.GetOptions{})
	if err != nil {
		_,err := cli.CoreV1().Namespaces().Create(&v1.Namespace{
			ObjectMeta:metav1.ObjectMeta{
				Name:ns,
			},
		})
	    if err != nil {
			fmt.Println("create ns err",err)
			return
		}
    }
	return
}

func InitAll(){
	initNs()
}