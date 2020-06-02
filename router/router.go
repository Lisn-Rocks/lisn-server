package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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
					<form enctype="multipart/form-data" action="/upload"
						method="POST">
					<input type="password" name="password"
						placeholder="Password"> <br>
					<input type="file" name="album"> <br>
					<input type="submit" value="Upload">
					</form>
				</body>
			</html>`
			fmt.Fprint(w, site)
		},
	)

	root.Subrouter().Path("/upload").Methods(http.MethodPost).
		Handler(NewHandler(env,
			func(w http.ResponseWriter, r *http.Request, e *Env) (re error) {
				e.logr.Print("Incoming upload request")

				// Parse form and check authenticity.
				if err := r.ParseMultipartForm(config.MaxMemUploadSize); err != nil {
					e.logr.Println(err)
				}

				password := r.FormValue("password")

				// I know that hardcoding a hash is a crude way to check
				// identity but it has to stay until we can properly handle user
				// permissions.
				if util.Hash(password, config.Salt) != config.Hash {
					e.logr.Print("Authentication failed (hash did not match)")
					fmt.Fprint(w, "<h1>You aren't authorized to upload.</h1>")
					return
				}

				album, _, err := r.FormFile("album")
				if err != nil {
					e.logr.Println(err)
					fmt.Fprint(w, "<h1>Failed to retreive archive.</h1>")
					return
				}
				defer album.Close()

				// Save archive under random filename
				apath := path.Join(config.StorageFolder, util.RandName()+".zip")

				/*
					for _, err := os.Stat(apath); !os.IsNotExist(err); apath = path.Join(config.StorageFolder, util.RandName()+".zip") {
					}
				*/

				archive, _ := os.Create(apath)
				io.Copy(archive, album)
				archive.Close()
				defer os.Remove(apath)

				meta, err := util.ReadMeta(apath)
				if err != nil {
					e.logr.Println(err)
				} else {
					e.logr.Println(meta.Album)
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
