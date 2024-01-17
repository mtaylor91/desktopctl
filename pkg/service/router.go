package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func router(s *Service) *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/api/vms", s.listVMs).Methods("GET")
	api.HandleFunc("/api/vms/{name}", s.getVM).Methods("GET")
	api.HandleFunc("/api/vms/{name}/start", s.startVM).Methods("POST")
	api.HandleFunc("/api/vms/{name}/stop", s.stopVM).Methods("POST")
	api.Use(s.auth)

	r := mux.NewRouter()
	r.PathPrefix("/api").Handler(api)
	r.HandleFunc("/auth/callback", s.authCallback).Methods("GET")
	r.HandleFunc("/auth/login", s.authLogin).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(staticFS()))

	return r
}
