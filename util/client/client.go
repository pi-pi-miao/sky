package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	return clientset,nil
}