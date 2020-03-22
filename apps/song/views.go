package song

import (
    "net/http"
    "strings"
    "log"
    "io"
    "strconv"
    "database/sql"
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
    songid, err := strconv.Atoi(split[1])

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



// ServeRandom serves one random song when called.
func ServeRandom(
    w http.ResponseWriter, r *http.Request,
    db *sql.DB, logr *log.Logger,
) {
    // Count songs
    var n int
    row := db.QueryRow(`SELECT count(*) FROM songs;`)
    row.Scan(&n)


    songid := rand.Intn(n) + 1
    extension, err := getSongExtension(songid, db)

    if err != nil {
        logr.Print("Cannot serve random song due to database error")

        w.WriteHeader(http.StatusInternalServerError)
        io.WriteString(w, "500 internal server error")

        return
    }


    serveFileFromFolder(w, r, logr, config.SongsFolder, songid, extension)
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
    split := strings.Split( r.URL.String(), "/" )[1:]

    // Atoi used to prevent injections. Only numbers pass!
    id, err := strconv.Atoi(split[1])

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
    serveFileFromFolder(w, r, logr, config.AlbumsFolder, albumid, extension)
}
