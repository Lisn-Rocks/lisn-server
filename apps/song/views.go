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
	"database/sql"
	"errors"
	"math/rand"

    "github.com/sharpvik/Lisn/config"
)


// ServeByID function is used to serve songs using their databse ID.
//
// Example URL:
//
//     http://localhost:8000/song/42
//
func ServeByID(
    w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
    split := strings.Split( r.URL.String(), "/" )[1:]

	// Atoi used to prevent injections. Only numbers pass!
	id, err := strconv.Atoi(split[1])

	if err != nil {
		logr.Print("Cannot convert ID specified in URL to int")

		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "400 forbidden")

		return
	}


	extension, err := getSongExtension(id, db)

	if err != nil {
		logr.Print("Error occurred while trying to fetch song's extension")

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 song not found in database")

		return
	}


	serveAudioFile(w, r, logr, id, extension)
}


func ServeRandom(
	w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
	// Count songs
	var n int
	row := db.QueryRow(`SELECT count(*) FROM songs;`)
	row.Scan(&n)


	id := rand.Intn(n) + 1
	extension, err := getSongExtension(id, db)

	if err != nil {
		logr.Print("Cannot serve random song due to database error")

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "500 internal server error")

		return
	}
	

	serveAudioFile(w, r, logr, id, extension)
}



func getSongExtension(id int, db *sql.DB) (string, error) {
	var extension string

	row := db.QueryRow(`SELECT extension FROM songs WHERE songid=$1;`, id)

	switch err := row.Scan(&extension); err {
	case sql.ErrNoRows:
		return "",
			errors.New("Song with given ID does not exist in the database")

	case nil:
		return extension, nil

	default:
		return "", errors.New("Unknown error occurred")
	}
}


func serveAudioFile(
	w http.ResponseWriter, r *http.Request, logr *log.Logger,
	id int, extension string,
) {
	filepath := path.Join(
		config.SongsFolder,
		fmt.Sprintf("%d%s", id, extension),
	)


	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		logr.Printf("Cannot serve. Song '%s' not found in filesystem", filepath)

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 file not found")

		return
	}


    logr.Printf("Serving song at %s", filepath)
    http.ServeFile(w, r, filepath)
}
