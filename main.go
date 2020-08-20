package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Lisn-Rocks/server/config"
	"github.com/Lisn-Rocks/server/dbi"
	"github.com/Lisn-Rocks/server/router"
)

func main() {
	logr := log.New(config.LogWriter, config.LogPrefix, log.Ltime)
	dbi := dbi.Init(logr)
	env := router.NewEnv(logr, dbi)
	config.Init()

	server := http.Server{
		Addr:     config.Port,
		Handler:  router.Init(env),
		ErrorLog: logr,
	}

	logr.Printf("serving at port %s", config.Port)
	logr.Fatalln(server.ListenAndServe())
}
