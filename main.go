package main

import (
    "database/sql"
    "os"
    "net/http"
    "log"
    "strings"

    _ "github.com/mattn/go-sqlite3"

    "github.com/sharpvik/Lisn/config"
    "github.com/sharpvik/Lisn/apps/index"
    "github.com/sharpvik/Lisn/apps/song"
    "github.com/sharpvik/Lisn/apps/favicon"
    "github.com/sharpvik/Lisn/apps/public"
)



var logr *log.Logger

var mux *http.ServeMux

var db *sql.DB



func main() {
    initRequired := config.InitRequired()

    db, _ = sql.Open("sqlite3", config.DatabaseFile)

    if initRequired {
        initDB(db)
        insertSongs(db)
    }


    logr = log.New(os.Stdout, "", log.Ltime)


    server := http.Server{
        Addr:       config.Port,
        Handler:    &mainHandler{},
        ErrorLog:   logr,
    }


    // mux is not used as there are no apps that use simple URL patterns.
    // However, it may be useful in the future, so we keep it here.
    mux = http.NewServeMux()


    logr.Printf("Serving at localhost%s", config.Port)
    server.ListenAndServe()
}



type mainHandler struct {}

// ServeHTTP function is the entry point for server's routing mechanisms.
// It is used to delegate request to a proper handler function.
func (*mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    url := stripParams( r.URL.String() )
    logr.Printf("URL: %s", url)

    var app string

    if url == "/" {
        app = "index"

    } else {
        split := strings.Split(url, "/")[1:]
    
        // First string in the split must name the app for which this request is
        // being made. That helps keep app routing at O(1).
        app = split[0]
    }

    logr.Printf("App: %s", app)

    switch app {
    case "index":
        index.Serve(w, r, db)

    case "song":
        song.ServeByID(w, r, db, logr)

    case "public":
        public.ServeFile(w, r, logr)

    default:
        mux.ServeHTTP(w, r)
    }
}



func initDB(db *sql.DB) {
    stmt, _ := db.Prepare("CREATE TABLE IF NOT EXISTS songs (id INTEGER PRIMARY KEY, title TEXT, duration INTEGER, genre TEXT, artist TEXT, album TEXT NULL)")
    stmt.Exec()
}

func insertSongs(db *sql.DB) {
    stmt, _ := db.Prepare("INSERT INTO songs (title, duration, genre, artist, album) VALUES (?, ?, ?, ?, ?)")
    stmt.Exec("Another One Bites the Dust", 222, "Classic Rock", "Queen", "The Game")
    stmt.Exec("Don't Stop Me Now", 217, "Classic Rock", "Queen", "Jazz")
    stmt.Exec("I Want To Break Free", 271, "Classic Rock", "Queen", "The Works")
    stmt.Exec("Somebody To Love", 309, "Rock", "Queen", "A Day at the Races")
}


func stripParams(url string) string {
    questionmarkIndex := strings.IndexByte(url, '?')

    if questionmarkIndex != -1 {
        return url[:questionmarkIndex]
    }

    return url
}
