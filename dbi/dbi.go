package dbi

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/sharpvik/lisn-server/config"
)

// DBI stands for DataBase Interface. It is a struct that wraps around *sql.DB
// and provides essential methods to interface with Lisn databases.
type DBI struct {
	*sql.DB
	logr *log.Logger
}

// Init function tries to connect to the database and ping it. We can only hope
// that all goes well.
func Init(logr *log.Logger) *DBI {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBhost, config.DBport, config.DBuser,
		config.DBpassword, config.DBname,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		logr.Print("database arguments are invalid")
		logr.Fatalln(err)
	}

	err = db.Ping()

	if err != nil {
		logr.Print("can't connect to the database")
		logr.Fatalln(err)
	}

	return &DBI{db, logr}
}

// Close gives acces to the `db.Close` function from without the package. It's
// used in the main function.
func (dbi *DBI) Close() {
	dbi.Close()
}

// CoverExtension returns song's albums cover file's extension as it is on
// the server's file system if it exists, otherwise returns non-nil error.
func (dbi *DBI) CoverExtension(albumid int) (extension string, err error) {
	row := dbi.QueryRow(
		`SELECT extension FROM albums WHERE albumid=$1;`, albumid,
	)
	err = row.Scan(&extension)
	return
}

// ArtistName returns artist name based on artistid if such id exists in the
// database otherwise returns non-nil error.
func (dbi *DBI) ArtistName(artistid int) (name string, err error) {
	row := dbi.QueryRow(`SELECT name FROM artists WHERE artistid=$1;`, artistid)
	err = row.Scan(&name)
	return
}

// Exists tells you whether song with given id exists in the database.
func (dbi *DBI) Exists(songid int) (answer bool) {
	var name string
	row := dbi.QueryRow(`SELECT name FROM songs WHERE songid=$1;`, songid)
	err := row.Scan(&name)
	return err == nil
}

// Extension returns song file's extension if said song exists, otherwise it
// returns a non-nil error.
func (dbi *DBI) Extension(songid int) (extension string, err error) {
	row := dbi.QueryRow(`SELECT extension FROM songs WHERE songid=$1;`, songid)
	err = row.Scan(&extension)
	return
}

// AlbumName returns song's album name if that song exists, otherwise returns a
// non-nil error.
func (dbi *DBI) AlbumName(songid int) (name string, err error) {
	albumid, err := dbi.albumID(songid)

	if err != nil {
		return
	}

	row := dbi.QueryRow(`SELECT name FROM albums WHERE albumid=$1;`, albumid)
	err = row.Scan(&name)
	return
}

func (dbi *DBI) albumID(songid int) (albumid int, err error) {
	row := dbi.QueryRow(`SELECT albumid FROM songs WHERE songid=$1;`, songid)
	err = row.Scan(&albumid)
	return
}
