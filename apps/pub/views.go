package pub

import (
    "log"
    "net/http"
    "path"
    "os"

    "github.com/sharpvik/Lisn/config"
    "github.com/sharpvik/Lisn/util"
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
    url := r.URL.String()
    fullPath := path.Join(config.RootFolder, url)

    if _, err := os.Stat(fullPath); os.IsNotExist(err) {
        util.FailWithCode(w, r, http.StatusNotFound, logr)
        return
    }

    logr.Printf("Serving file <Root>%s", url)
    http.ServeFile(w, r, fullPath)
}
