package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/sharpvik/lisn-server/config"
	"github.com/sharpvik/lisn-server/router"
)

// main function is very simple. It merely opens and connects to the database,
// initializes logger and starts serving, invoking `server.ListenAndServe()`.
func main() {
	logr := log.New(config.LogWriter, config.LogPrefix, log.Ltime)
	//dbi := dbi.Init(logr)

	server := http.Server{
		Addr:     config.Port,
		Handler:  router.Init(logr),
		ErrorLog: logr,
	}

	logr.Printf("Serving at port %s", config.Port)
	server.ListenAndServe()
}

/*
type mainHandler struct{}

// ServeHTTP function is the entry point for server's routing mechanisms.
// It is used to delegate request to a proper handler function.
func (*mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	url := r.URL.String()
	logr.Printf("URL: %s", url)

	// The following if statements are used to redirect common URL requests
	// towards a proper internal server app. For example, empty URL must refer
	// to Lisn's default index.html page.
	//
	// Later, this logic may be somehow substituted to accomodate for the user
	// authentication practices. In that sense, authenticated users must be
	// redirected into the app, while those who aren't authenticated will be
	// sent over to the registration (lisnup) page.
	if url == "/" {
		url = "/pub/lisn/index.html"
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	if url == "/favicon.ico" {
		url = "/pub/lisn/favicon.ico"
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	split := strings.Split(url, "/")[1:]

	if len(split) < 2 {
		util.FailWithCode(w, r, http.StatusNotFound, logr)
		return
	}

	// First string in the split must name the server mode. There are only two
	// of them:
	//
	//     * api -- API mode that's used by developers
	//     * pub -- Public mode that's used to serve files from the 'pub' folder
	//
	mode := split[0]

	// Second string in the split must specify the app you want to use from the
	// given server mode. For example, api server mode has an app called 'cover'
	// that serves album's cover image for a given song ID.
	app := split[1]

	logr.Printf("Route: %s -> %s", mode, app)

	if mode == "api" {
		switch app {
		case "song":
			api.ServeByID(w, r, db, logr)

		case "random":
			api.ServeRandom(w, r, db, logr)

		case "cover":
			api.ServeCover(w, r, db, logr)

		case "covermin":
			api.ServeCoverMin(w, r, db, logr)

		case "info":
			api.ServeJSON(w, r, db, logr)

		default:
			util.FailWithCode(w, r, http.StatusNotFound, logr)
		}
	} else if mode == "pub" {
		pub.ServeFile(w, r, logr)
	} else {
		util.FailWithCode(w, r, http.StatusNotFound, logr)
	}
}
*/
