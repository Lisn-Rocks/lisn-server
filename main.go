package main

import (
	"database/sql"
    "os"
    "net/http"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/sharpvik/lisn-backend/config"
	"github.com/sharpvik/lisn-backend/apps/index"
)



var logr *log.Logger

var mux *http.ServeMux

var db *sql.DB



func main() {
	db, _ = sql.Open("sqlite3", "./songs.db")

	if config.InitRequired {
		initDB(db)
		insertSongs(db)
	}


	logr = log.New(os.Stdout, "", log.Ltime)


	server := http.Server{
		Addr:		config.Port,
		Handler:	&mainHandler{},
		ErrorLog:	logr,
	}


	mux = http.NewServeMux()
	mux.HandleFunc("/id", serveSongByID)


	logr.Printf("Serving at localhost%s", config.Port)
	server.ListenAndServe()
}



type mainHandler struct {}

// ServeHTTP function is the entry point for server's routing mechanisms.
// It uses mux to delegate request to a proper handler function.
func (*mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logr.Printf( "URL: %s", r.URL.String() )

	switch r.URL.String() {
	case "/":
		index.Serve(w, r, db)

	default:
		mux.ServeHTTP(w, r)
	}
}



func serveSongByID(w http.ResponseWriter, r *http.Request) {
	id := "1" // in production id must be read from request body (type string)
	path := fmt.Sprintf("store/%s.mp3", id)
	http.ServeFile(w, r, path)
}



func initDB(db *sql.DB) {
	stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS songs (id INTEGER PRIMARY KEY, title TEXT, duration FLOAT, genre TEXT, author TEXT, album TEXT NULL)")
	stmt.Exec()
}

func insertSongs(db *sql.DB) {
	stmt, _ := db.Prepare("INSERT INTO songs (title, duration, genre, author, album) VALUES (?, ?, ?, ?, ?)")
	stmt.Exec("Another One Bites the Dust", 3.7, "Classic Rock", "Queen", "The Game")
	stmt.Exec("Don't Stop Me Now", 3.617, "Classic Rock", "Queen", "Jazz")
	stmt.Exec("I Want To Break Free", 4.517, "Classic Rock", "Queen", "The Works")
	stmt.Exec("Somebody To Love", 5.15, "Rock", "Queen", "A Day at the Races")
}
