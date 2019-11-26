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
}
