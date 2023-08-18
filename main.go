package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/mtaylor91/desktopctl/pkg/kubevirt/client/versioned"
	"github.com/mtaylor91/desktopctl/pkg/service"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const inClusterToken = "/var/run/secrets/kubernetes.io/serviceaccount/token"

func getKubeConfig() (*rest.Config, error) {
	var kubeconfig *string

	if _, err := os.Stat(inClusterToken); err == nil {
		return rest.InClusterConfig()
	} else if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "",
			"absolute path to the kubeconfig file")
	}

	flag.Parse()

	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}

func main() {
	config, err := getKubeConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	err = service.New(":8080", "kubevirt", clientset).Start()
	if err != nil {
		panic(err)
	}
}
