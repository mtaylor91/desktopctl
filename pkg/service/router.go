package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router(s *Service) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/vms", s.listVMs).Methods("GET")
	r.HandleFunc("/api/vms/{name}", s.getVM).Methods("GET")
	r.HandleFunc("/api/vms/{name}/start", s.startVM).Methods("POST")
	r.HandleFunc("/api/vms/{name}/stop", s.stopVM).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(staticFS()))
	return r
}
