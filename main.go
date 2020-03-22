package main

import (
    "database/sql"
    "os"
    "net/http"
    "log"
    "strings"
    "fmt"

    _ "github.com/lib/pq"

    "github.com/sharpvik/Lisn/config"
    "github.com/sharpvik/Lisn/apps/song"
    /*
    "github.com/sharpvik/Lisn/apps/songinfo"
    "github.com/sharpvik/Lisn/apps/public"
    */
)



var logr *log.Logger

var mux *http.ServeMux

var db *sql.DB



func main() {
    var err error // declaring it here so that global db is used on sql.Open

    psqlInfo := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.DBhost, config.DBport, config.DBuser,
        config.DBpassword, config.DBname,
    )

    db, err = sql.Open("postgres", psqlInfo)

    if err != nil {
        panic("Database arguments are invalid")
    }

    defer db.Close()

    err = db.Ping()

    if err != nil {
        panic("Can't connect to the database")
    }


    logr = log.New(os.Stdout, "", log.Ltime)


    server := http.Server{
        Addr:       config.Port,
        Handler:    &mainHandler{},
        ErrorLog:   logr,
    }


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


    if url == "/" {
        url = "/public/index.html"
    }

    split := strings.Split(url, "/")[1:]

    // First string in the split must name the app for which this request is
    // being made. That helps keep app routing at O(1).
    app := split[0]


    logr.Printf("App: %s", app)

    switch app {
    case "song":
        song.ServeByID(w, r, db, logr)

    case "random":
        song.ServeRandom(w, r, db, logr)
/*
    case "songinfo":
        songinfo.ServeJSON(w, r, db, logr)

    case "public":
        public.ServeFile(w, r, logr)
*/
    default:
        mux.ServeHTTP(w, r)
    }
}


/* This function is not used for now but may be useful later
func stripParams(url string) string {
    questionmarkIndex := strings.IndexByte(url, '?')

    if questionmarkIndex != -1 {
        return url[:questionmarkIndex]
    }

    return url
}
*/
