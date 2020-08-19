package router

import (
	"log"
	"net/http"

	"github.com/Lisn-Rocks/server/dbi"
	"github.com/Lisn-Rocks/server/util"
)

// Env is a struct that carries environment data for request handlers.
type Env struct {
	logr *log.Logger
	dbi  *dbi.DBI
}

// NewEnv returns pointer to a newly created Env instance.
func NewEnv(logr *log.Logger, dbi *dbi.DBI) *Env {
	return &Env{logr, dbi}
}

// View is a function type used by the Handler to handle requests that need
// access to data stored in Env.
type View func(http.ResponseWriter, *http.Request, *Env) error

// Handler implements http.Handler interface. It is useful because it can carry
// environment data.
type Handler struct {
	E *Env
	F View
}

// NewHandler returns pointer to a newly created Handler instance.
func NewHandler(env *Env, fun View) *Handler {
	return &Handler{env, fun}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.F(w, r, h.E); err != nil {
		util.FailWithCode(w, r, http.StatusInternalServerError, h.E.logr)
	}
}
