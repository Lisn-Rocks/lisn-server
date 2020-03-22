package song

// This file contains package private functions.

import (
    "net/http"
    "log"
    "database/sql"
    "errors"
    "path"
    "fmt"
    "os"
    "io"

    "github.com/sharpvik/Lisn/config"
)

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
