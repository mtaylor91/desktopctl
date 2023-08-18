package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/mtaylor91/desktopctl/pkg/kubevirt/client/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "",
			"absolute path to the kubeconfig file")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	vms, err := clientset.KubevirtV1().
		VirtualMachines("kubevirt").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, instance := range vms.Items {
		fmt.Printf("VM: %s (%s)\n",
			instance.Name, instance.Status.PrintableStatus)
	}
}
