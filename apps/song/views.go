package song

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"path"
	"os"
	"math/rand"
	"net/http"

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
	songid, err := parseIDFromURL(r)

	if err != nil {
		logr.Print("Cannot convert song id specified in URL to int")

		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "400 forbidden")

		return
	}

	if !songExists(songid, db) {
		logr.Print("Song with given ID does not exist in the database")

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 song not found in database")

		return
	}

	extension, _ := getSongExtension(songid, db)
	serveFileFromFolder(w, r, logr, config.SongsFolder, songid, extension)
}

// ServeRandom serves one random song when called by generating a random song ID
// and redirecting to the URL that invokes ServeByID.
func ServeRandom(
	w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
	// Count songs
	var n int
	row := db.QueryRow(`SELECT count(*) FROM songs;`)
	row.Scan(&n)

    songid := rand.Intn(n) + 1
    
	http.Redirect(
		w, r,
		fmt.Sprintf("/info/%d", songid),
		http.StatusSeeOther,
	)
}

// ServeCover serves song's cover image by its database ID.
//
// Example URL:
//
//     http://localhost:8000/cover/42
//
func ServeCover(
	w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
	id, err := parseIDFromURL(r)

	if err != nil {
		logr.Print("Cannot convert ID specified in URL to int")

		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "400 forbidden")

		return
	}

	if !songExists(id, db) {
		logr.Print("Song with given ID does not exist in the database")

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 song not found in database")

		return
	}

	albumid, _ := getAlbumID(id, db)
	extension, _ := getAlbumExtension(albumid, db)

	filepath := path.Join(
		config.AlbumsFolder,
		fmt.Sprintf("%d-min%s", albumid, extension),
	)

    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        logr.Printf("Cannot serve. File '%s' not found in filesystem", filepath)

        w.WriteHeader(http.StatusNotFound)
        io.WriteString(w, "404 file not found")

        return
    }

    logr.Printf("Serving file %s", filepath)
    http.ServeFile(w, r, filepath)
}

// ServeJSON function serves song data in JSON file. Song's database ID must be
// specified in request URL.
//
// Example URL:
//
//     http://localhost:8000/info/59
//
func ServeJSON(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	logr *log.Logger,
) {
	songid, err := parseIDFromURL(r)

	if err != nil {
		logr.Print("Cannot convert ID specified in URL to int")

		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, "400 forbidden")

		return
	}

	if !songExists(songid, db) {
		logr.Print("Song with given ID does not exist in the database")

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404 song not found in database")

		return
	}

	row := db.QueryRow(
		`SELECT name, artistid, genre, albumid FROM songs WHERE songid=$1;`,
		songid,
	)

	var artistid, albumid int
	var name, genre string

	row.Scan(&name, &artistid, &genre, &albumid)

	artist, err := getArtistName(artistid, db)
	album, err := getAlbumName(songid, db)

	if err != nil {
		logr.Print("Failed to obtain artist or album name")

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "500 internal server error")

		return
	}

	asong := song{
		songid,
		name,
		artist,
		genre,
		album,
	}

	jsn, err := json.Marshal(asong)

	if err != nil {
		logr.Printf("Error: %s", err)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	// Remove header below before building for production.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
