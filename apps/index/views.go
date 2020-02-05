package index

import (
	"net/http"
)


// Redirect fucntion redirects user towards index.html file.
func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/public/index.html", http.StatusMovedPermanently)
}
