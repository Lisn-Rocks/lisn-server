package song

import (
    "net/http"
    "strings"
    "log"
    "io"
    "strconv"
    "database/sql"
    "math/rand"
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


// ServeRandom serves one random song when called.
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
