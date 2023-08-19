package service

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed client/*
var clientFS embed.FS

func staticFS() http.FileSystem {
	subFS, err := fs.Sub(clientFS, "client")
	if err != nil {
		panic(err)
	}

	return http.FS(subFS)
}
