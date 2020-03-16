package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"flag"
)

var (
	Client *kubernetes.Clientset
)

// todo add config to cli
func GetClient()(*kubernetes.Clientset,error){
	kubeconfig := flag.String("kubeconfig", "/root/config", "")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil,err
	}
	Client = clientset
	return Client,nil
}
