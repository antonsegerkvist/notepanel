package notepanel

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/notepanel/api"
	"github.com/notepanel/web"
)

//
// Run mounts all the server endpoints and boots the server.
//
func Run() {

	router := mux.NewRouter()

	api.Mount(router)
	web.Mount(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())

}
