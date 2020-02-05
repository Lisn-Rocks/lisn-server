package songinfo

import (
	"net/http"
	"fmt"
	"database/sql"
	"strings"
	"log"
	"encoding/json"
	"strconv"
)


// ServeJSON function serves song data in JSON file. Song's database ID must be
// specified in request URL. 
//
// Example URL:
//
//     http://localhost:8000/songinfo/42
//
func ServeJSON(
	w http.ResponseWriter,
    r *http.Request,
    db *sql.DB,
	logr *log.Logger,
) {
	split := strings.Split( r.URL.String(), "/" )[1:]

	// Atoi used to prevent SQL injections. Only numbers pass!
	id, err := strconv.Atoi(split[1])

	if err != nil {
		logr.Print("Cannot convert ID specified in URL to int")
		return
	}


	rows, _ := db.Query(
		fmt.Sprintf("SELECT title, duration, genre, artist, album FROM songs WHERE id=%d", id),
	)

	var duration int
	var title, genre, artist, album string

	rows.Next()
	rows.Scan(&title, &duration, &genre, &artist, &album)


	asong := song{id, title, duration, genre, artist, album}
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
