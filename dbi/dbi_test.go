package dbi

import (
	"log"
	"testing"

	_ "github.com/lib/pq"

	"github.com/sharpvik/lisn-server/config"
)

func TestArtistID(t *testing.T) {
	logr := log.New(config.LogWriter, config.LogPrefix, log.Ltime)
	dbi := Init(logr)

	artistid := dbi.ArtistID("Some Artist")
	if artistid != 0 {
		t.Errorf("artistid expected 0; got %d", artistid)
	}

	artistid = dbi.ArtistID("Queen")
	if artistid != -1 {
		t.Errorf("artistid expected -1; got %d", artistid)
	}

	artistid = dbi.ArtistID("Michael Jackson")
	if artistid != 3 {
		t.Errorf("artistid expected 3; got %d", artistid)
	}
}
