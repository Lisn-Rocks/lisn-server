package router

import (
	"fmt"
	"net/http"

	"github.com/sharpvik/lisn-server/config"
	"github.com/sharpvik/mux"
)

// Init returns pointer to the router.
func Init(env *Env) *mux.Router {
	root := mux.New()

	root.Subrouter().PathPrefix("/pub/").Methods(http.MethodGet).Handler(
		http.FileServer(http.Dir(config.PublicFolder)))

	root.Subrouter().Path("/upload").Methods(http.MethodPost).
		Handler(NewHandler(env, processUpload)) // views.go > processUpload

	api := root.Subrouter().PathPrefix("/api")
	initAPI(api)

	root.Subrouter().Path("/favicon.ico").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			url := "/pub/lisn/favicon.ico"
			http.Redirect(w, r, url, http.StatusSeeOther)
		},
	)

	// The forward slash path matches any request. Therefore, any requests that
	// don't match above subrouters will be handled by this function. They will
	// be redirected to "/pub/lisn".
	root.Subrouter().Path("/").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			url := "/pub/lisn/"
			http.Redirect(w, r, url, http.StatusSeeOther)
		},
	)

	return root
}

func initAPI(api *mux.Router) {
	api.Subrouter().PathPrefix("/album").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "album site")
		},
	)
	api.Subrouter().PathPrefix("/discover").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "discover site")
		},
	)
	api.Subrouter().PathPrefix("/song").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "song site")
		},
	)
}
