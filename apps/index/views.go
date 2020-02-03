package index

import (
	"net/http"
	"io"
    "database/sql"
)


// Serve function is used to serve the main index of songs to the user.
func Serve(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	io.WriteString(w, "Lisn Index Page");
}
