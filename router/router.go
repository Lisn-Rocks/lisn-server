package router

import (
	"log"
	"net/http"

	"github.com/sharpvik/lisn-server/config"
	"github.com/sharpvik/mux"
)

// Init returns pointer to the router.
func Init(logr *log.Logger) *mux.Router {
	root := mux.New()

	root.Subrouter().PathPrefix("/pub/").Handler(
		http.FileServer(http.Dir(config.PublicFolder)),
	)

	root.Subrouter().Path("/").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			url := "/pub/lisn/"
			http.Redirect(w, r, url, http.StatusSeeOther)
		},
	)

	return root
}
