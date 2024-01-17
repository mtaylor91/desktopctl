package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"

	"github.com/coreos/go-oidc"
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
	ctx := context.Background()

	config, err := getKubeConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		namespace = "kubevirt"
	}

	oidcClientID := os.Getenv("OIDC_CLIENT_ID")
	oidcClientSecret := os.Getenv("OIDC_CLIENT_SECRET")
	oidcIssuerURL := os.Getenv("OIDC_ISSUER_URL")
	oidcRedirectURL := os.Getenv("OIDC_REDIRECT_URL")
	oidcProvider, err := oidc.NewProvider(ctx, oidcIssuerURL)
	if err != nil {
		panic(err)
	}

	err = service.New(
		":8080",
		namespace,
		oidcClientID,
		oidcClientSecret,
		oidcRedirectURL,
		oidcProvider,
		clientset,
	).Start()
	if err != nil {
		panic(err)
	}
}
