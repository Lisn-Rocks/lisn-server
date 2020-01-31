package favicon

import (
	"net/http"
	"path"

	"github.com/sharpvik/Lisn/config"
)


// Serve is a simple view that sends site's favicon over.
func Serve(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(
		w, r,
		path.Join(config.StaticFolder, "favicon.gif"),
	)
}
