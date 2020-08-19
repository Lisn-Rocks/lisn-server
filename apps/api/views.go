package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"path"

	"github.com/Lisn-Rocks/server/config"
	"github.com/Lisn-Rocks/server/util"
)

// ServeByID function is used to serve songs using their databse ID.
//
// Example URL:
//
//     http://localhost:8000/api/song/42
//
func ServeByID(
	w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
	songid := 42

	if !songExists(songid, db) {
		logr.Print("song with given ID does not exist in the database")
		util.FailWithCode(w, r, http.StatusNotFound, logr)
		return
	}

	extension, _ := getSongExtension(songid, db)
	filepath := path.Join(
		config.SongsFolder,
		fmt.Sprintf("%d%s", songid, extension),
	)
	serveFile(w, r, logr, filepath)
}

// ServeRandom serves one random song when called by generating a random song ID
// and redirecting to the URL that invokes ServeByID.
//
// Example URL:
//
//     http://localhost:8000/api/random
//
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
		fmt.Sprintf("/api/info/%d", songid),
		http.StatusSeeOther,
	)
}

// ServeCover serves song's cover image by its database ID.
//
// Example URL:
//
//     http://localhost:8000/api/cover/42
//
func ServeCover(
	w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
	id := 42

	if !songExists(id, db) {
		logr.Print("song with given ID does not exist in the database")
		util.FailWithCode(w, r, http.StatusNotFound, logr)
		return
	}

	albumid, _ := getAlbumID(id, db)
	extension, _ := getAlbumExtension(albumid, db)

	filepath := path.Join(
		config.AlbumsFolder,
		fmt.Sprintf("%d%s", albumid, extension),
	)
	serveFile(w, r, logr, filepath)
}

// ServeCoverMin serves song's minimized cover image by its database ID.
//
// Example URL:
//
//     http://localhost:8000/api/covermin/42
//
func ServeCoverMin(
	w http.ResponseWriter, r *http.Request,
	db *sql.DB, logr *log.Logger,
) {
	id := 42

	if !songExists(id, db) {
		logr.Print("Song with given ID does not exist in the database")
		util.FailWithCode(w, r, http.StatusNotFound, logr)
		return
	}

	albumid, _ := getAlbumID(id, db)
	extension, _ := getAlbumExtension(albumid, db)

	filepath := path.Join(
		config.AlbumsFolder,
		fmt.Sprintf("%d-min%s", albumid, extension),
	)
	serveFile(w, r, logr, filepath)
}

// ServeJSON function serves song data in JSON file. Song's database ID must be
// specified in request URL.
//
// Example URL:
//
//     http://localhost:8000/api/info/59
//
func ServeJSON(
	w http.ResponseWriter,
	r *http.Request,
	db *sql.DB,
	logr *log.Logger,
) {
	songid := 42

	if !songExists(songid, db) {
		logr.Print("Song with given ID does not exist in the database")
		util.FailWithCode(w, r, http.StatusNotFound, logr)
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
		util.FailWithCode(w, r, http.StatusInternalServerError, logr)
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

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}
