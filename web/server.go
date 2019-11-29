package web

import "github.com/gorilla/mux"

import "net/http"

//
// Mount mounts the static file server to the router.
//
func Mount(router *mux.Router) {
	fs := http.FileServer(http.Dir("frontend/dist"))
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", fs))
}
