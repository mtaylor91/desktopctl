package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mtaylor91/desktopctl/pkg/kubevirt/client/versioned"
)

type Service struct {
	client    *versioned.Clientset
	namespace string
	router    *mux.Router
	server    *http.Server
}

func New(addr, namespace string, client *versioned.Clientset) *Service {
	s := &Service{client: client, namespace: namespace}
	s.router = router(s)
	s.server = &http.Server{Addr: addr, Handler: s.router}
	return s
}
