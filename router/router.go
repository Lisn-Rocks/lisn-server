package router

import (
	"fmt"
	"net/http"

	"github.com/sharpvik/lisn-server/config"
	"github.com/sharpvik/lisn-server/util"
	"github.com/sharpvik/mux"
)

// Init returns pointer to the router.
func Init(env *Env) *mux.Router {
	root := mux.New()

	root.Subrouter().PathPrefix("/pub/").Methods(http.MethodGet).Handler(
		http.FileServer(http.Dir(config.PublicFolder)),
	)

	initUpload(root, env)

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

func initUpload(root *mux.Router, env *Env) {
	root.Subrouter().Path("/upload").Methods(http.MethodGet).HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			site := `
			<html>
				<head><title>Album Upload</title></head>
				<body>
					<form action="http://localhost:8000/upload" method="POST">
					<input type="password" name="password"
						placeholder="Password"> <br>
					<input type="file" name="folder" webkitdirectory directory
						multiple> <br>
					<input type="submit" value="Upload">
					</form>
				</body>
			</html>`
			fmt.Fprint(w, site)
		},
	)

	root.Subrouter().
		Path("/upload").
		Methods(http.MethodPost).
		Handler(NewHandler(env,
			func(w http.ResponseWriter, r *http.Request, e *Env) (re error) {
				e.logr.Print("Incoming upload request")

				// Parse form and check authenticity.
				if err := r.ParseMultipartForm(0); err != nil {
					e.logr.Println(err)
				}

				// I know that hardcoding a hash is a crude way to check
				// identity but it has to stay until we have a user tracker that
				// can properly handle user permissions.
				hash := `SOU9h8MjuADDO+tx5pQqt+GjMmI0RTyN1CiZXGSXQt1XVF117WKrprphqT03Obb0iwlf2DBBFXIp7qUZqCooPA==`
				salt := `N4M'R-d#23X@70[jZZ&r4#//+*E1W7,[a`
				password := r.FormValue("password")

				if util.Hash(password, salt) != hash {
					e.logr.Print("Authentication failed (hash did not match)")
					fmt.Fprint(w, "<h1>You aren't authorized to upload.</h1>")
					return
				}

				fmt.Fprint(w, "<h1>Your contribution is greatly appreciated!</h1>")
				return
			},
		))
}

func initAPI(api *mux.Router) {
	api.Subrouter().PathPrefix("/album").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "album site")
		},
	)
	api.Subrouter().PathPrefix("/random").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "random site")
		},
	)
	api.Subrouter().PathPrefix("/song").HandleFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "song site")
		},
	)
}
