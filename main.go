package main

import (
	"fmt"
	"net/http"

	"github.com/sharpvik/log-go"
)

const port = ":8080"

func init() {
	log.SetLevel(log.LevelDebug)
}

func main() {
	server := http.Server{
		Addr:     port,
		Handler:  &handler{},
	}

	log.Debug("serving at port %s", port)
	log.Fatal(server.ListenAndServe().Error())
}

type handler struct {}

func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ALL GOOD")
}
