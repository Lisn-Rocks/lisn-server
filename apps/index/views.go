package index

import (
	"net/http"
)


// Serve function is used to serve the main index of songs to the user.
func Serve(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/public/index.html", http.StatusMovedPermanently)
}
