package client

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClient(configPath string) (*kubernetes.Clientset, error) {
	var (
		kubeconfig *string
		err        error
	)
	kubeconfig = flag.String("kubeconfig", configPath, "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	Client, err := kubernetes.NewForConfig(config)
	return Client, err
}
