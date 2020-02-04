package public

import (
    "log"
    "net/http"
    "path"
	"os"
	"io"

    "github.com/sharpvik/Lisn/config"
)


// ServeFile function is a wrap-around that allows us to safely serve static
// files from the PublicFolder (specified in package config).
//
// ServeFile relies on URL param called 'path' that specifies subpath within
// the PublicFodler. Example URL:
//
//     http://localhost:8000/public/favicon.ico
//
func ServeFile(w http.ResponseWriter, r *http.Request, logr *log.Logger) {
	apath := r.URL.String()
    fullPath := path.Join(config.RootFolder, apath)

    if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		logr.Printf("Cannot serve. Path '%s' not found", fullPath)

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 file not found")

		return
    }

    http.ServeFile(w, r, fullPath)
}