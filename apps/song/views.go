package song

import (
    "net/http"
    "fmt"
    "database/sql"
	"path"
	"strings"
	"log"

    "github.com/sharpvik/Lisn/config"
)


// ServeByID function is used to serve songs using their databse ID.
func ServeByID(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	logr *log.Logger,
) {
	split := strings.Split( r.URL.String(), "/" )[1:]
	id := split[1]
	
	filepath := path.Join(
		config.StoreFolder,
		fmt.Sprintf("%s.mp3", id),
	)

	logr.Printf("Serving song at %s", filepath)

    http.ServeFile(w, r, filepath)
}
