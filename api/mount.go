package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//
// Mount mounts the api endpoints to the router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/api/note/list", handleNoteListGet).Methods(http.MethodGet)
	router.HandleFunc("/api/note/{id:[0-9]+}", handleNoteIDGet).Methods(http.MethodGet)
	router.HandleFunc("/api/note/{id:[0-9]+}", handleNoteIDPatch).Methods(http.MethodPatch)
	router.HandleFunc("/api/note", handleNotePost).Methods(http.MethodPost)
}
