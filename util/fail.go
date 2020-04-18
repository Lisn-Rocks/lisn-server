package util

import (
	"net/http"
	"log"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sharpvik/Lisn/config"
)

// FailWithCode is used to appropriately respond to user in an event of failure.
// It uses status parameter to figure out what exactly to send back.
func FailWithCode(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	logr *log.Logger,
) {
	fullPath := path.Join( config.FailFolder, fmt.Sprintf("%d.html", status) )

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
        w.WriteHeader(status)
        io.WriteString( w, fmt.Sprintf("FAIL #%d", status) )
        return
    }

	logr.Printf("Response: FAIL #%d", status)
    http.ServeFile(w, r, fullPath)
}