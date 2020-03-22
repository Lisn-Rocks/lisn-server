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
)



func songExists(songid int, db *sql.DB) bool {
    var name string
    row := db.QueryRow(`SELECT name FROM songs WHERE songid=$1;`, songid)
    err := row.Scan(&name)
    return err == nil
}



func getSongExtension(songid int, db *sql.DB) (string, error) {
    var extension string
    row := db.QueryRow(`SELECT extension FROM songs WHERE songid=$1;`, songid)
    err := row.Scan(&extension)

    if err != nil {
        return "", errors.New("Unexpected error occurred")
    }

    return extension, nil
}



func getAlbumID(songid int, db *sql.DB) (int, error) {
    var albumid int
    row := db.QueryRow(`SELECT albumid FROM songs WHERE songid=$1;`, songid)
    err := row.Scan(&albumid)

    if err != nil {
        return 0, errors.New("Unexpected error occurred")
    }

    return albumid, nil
}



func getAlbumExtension(albumid int, db *sql.DB) (string, error) {
    var extension string
    row := db.QueryRow(
        `SELECT extension FROM albums WHERE albumid=$1;`, albumid,
    )
    err := row.Scan(&extension);

    if err != nil {
        return "", errors.New("Unexpected error occurred")
    }

    return extension, nil
}



func serveFileFromFolder(
    w http.ResponseWriter, r *http.Request, logr *log.Logger,
    folder string, songid int, extension string,
) {
    filepath := path.Join( folder, fmt.Sprintf("%d%s", songid, extension) )


    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        logr.Printf("Cannot serve. File '%s' not found in filesystem", filepath)

        w.WriteHeader(http.StatusNotFound)
        io.WriteString(w, "404 file not found")

        return
    }


    logr.Printf("Serving file %s", filepath)
    http.ServeFile(w, r, filepath)
}
