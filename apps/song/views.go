package song

import (
	"net/http"
	"fmt"
	"path"
	"strings"
	"log"
	"os"
	"io"
	"strconv"

    "github.com/sharpvik/Lisn/config"
)


// ServeByID function is used to serve songs using their databse ID.
//
// Example URL:
//
//     http://localhost:8000/song/42
//
func ServeByID(
    w http.ResponseWriter,
    r *http.Request,
    logr *log.Logger,
) {
    split := strings.Split( r.URL.String(), "/" )[1:]
	
	// Atoi used to prevent SQL injections. Only numbers pass!
	id, err := strconv.Atoi(split[1])

	if err != nil {
		logr.Print("Cannot convert ID specified in URL to int")

		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "400 forbidden")

		return
	}

    
	filepath := path.Join(
		config.SongsAudioFolder,
		fmt.Sprintf("%d.mp3", id),
	)
	

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		logr.Printf("Cannot serve. Path '%s' not found", filepath)

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 file not found")

		return
	}


    logr.Printf("Serving song at %s", filepath)
    http.ServeFile(w, r, filepath)
}
