package dbi

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Lisn-Rocks/server/config"
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

// ArtistID returns artistid from the database. If artist with given name does
// not exist ArtistID returns 0, if there are two (or more) of such artists,
// ArtistID returns -1.
func (dbi *DBI) ArtistID(artist string) int {
	var n int
	dbi.QueryRow(`SELECT COUNT(*) FROM artists WHERE artist=$1`, artist).
		Scan(&n)

	switch n {
	case 0:
		return 0
	case 1:
		var artistid int
		dbi.QueryRow(`SELECT artistid FROM artists WHERE artist=$1`,
			artist).Scan(&artistid)
		return artistid
	default:
		return -1
	}
}

// AddArtist adds artist to the database and returns their artistid. In case of
// error, returns non-nil error.
func (dbi *DBI) AddArtist(artist string) (artistid int, err error) {
	err = dbi.QueryRow(`INSERT INTO artists (artist) VALUES ($1)
		RETURNING artistid`, artist).Scan(&artistid)
	return
}

// SearchAddArtist is a compex method that searches for given artist's artistid
// or -- if not found -- inserts one into the table and returns its artistid.
// Method returns non-nil error if there were more than two artists in the
// database or if there was an error adding new artist.
func (dbi *DBI) SearchAddArtist(artist string) (artistid int, err error) {
	artistid = dbi.ArtistID(artist)

	switch artistid {
	case -1:
		err = errors.New("multiple artists with the same name")
		return
	case 0:
		artistid, err = dbi.AddArtist(artist)
		return
	default:
		return
	}
}

// ArtistFromID returns artist's name given artistid. Returns an empty string if
// artistid out of bounds.
func (dbi *DBI) ArtistFromID(artistid int) (artist string) {
	dbi.QueryRow(`SELECT artist FROM artists WHERE artistid=$1`, artistid).
		Scan(&artist)
	return
}

// AlbumID returns albumid from the database. If album with given name does not
// exist AlbumID returns 0, if there are two (or more) of such albums, AlbumId
// returns -1.
func (dbi *DBI) AlbumID(album string, artistid int) int {
	var n int
	dbi.QueryRow(`SELECT COUNT(albumid) FROM albums
		WHERE album=$1 AND artistid=$2`, album, artistid).Scan(&n)

	switch n {
	case 0:
		return 0
	case 1:
		var albumid int
		dbi.QueryRow(
			`SELECT albumid FROM albums WHERE album=$1 AND artistid=$2`,
			album, artistid,
		).Scan(&albumid)
		return albumid
	default:
		return -1
	}
}

// AddAlbum adds album to the database and returns their albumid. In case of
// error, returns non-nil error.
func (dbi *DBI) AddAlbum(
	album string, artistid int, coverext string,
) (albumid int, err error) {
	err = dbi.QueryRow(`INSERT INTO albums (album, artistid, coverext)
		VALUES ($1, $2, $3) RETURNING albumid`, album, artistid, coverext).
		Scan(&albumid)
	return
}

// SearchAddAlbum is a compex method that searches for given album's albumid
// or -- if not found -- inserts one into the table and returns its albumid.
// Method returns -1 if there were more than two albums in the database.
func (dbi *DBI) SearchAddAlbum(
	album string, artistid int, coverext string,
) (albumid int, err error) {
	albumid = dbi.AlbumID(album, artistid)

	switch albumid {
	case -1:
		err = errors.New("multiple albums with the same name")
		return
	case 0:
		albumid, err = dbi.AddAlbum(album, artistid, coverext)
		return
	default:
		err = fmt.Errorf("album already exists: %s", album)
		return
	}
}

// CoverExtension returns song's albums cover file's extension as it is on
// the server's file system if it exists, otherwise returns non-nil error.
func (dbi *DBI) CoverExtension(albumid int) (extension string, err error) {
	err = dbi.QueryRow(`SELECT extension FROM albums WHERE albumid=$1;`,
		albumid).Scan(&extension)
	return
}

// ArtistName returns artist name based on artistid if such id exists in the
// database otherwise returns non-nil error.
func (dbi *DBI) ArtistName(artistid int) (name string, err error) {
	err = dbi.QueryRow(`SELECT name FROM artists WHERE artistid=$1;`, artistid).
		Scan(&name)
	return
}

// AddSong inserts a song record into the database and returns new song's songid
// for further use. Returns non-nil error if something went wrong.
func (dbi *DBI) AddSong(
	song string, albumid int, audioext string,
) (songid int, err error) {
	err = dbi.QueryRow(`INSERT INTO songs (song, albumid, audioext)
		VALUES ($1, $2, $3) RETURNING songid`, song, albumid, audioext).
		Scan(&songid)
	return
}

// SongExists tells you whether song with given id exists in the database.
func (dbi *DBI) SongExists(songid int) (answer bool) {
	var name string
	row := dbi.QueryRow(`SELECT name FROM songs WHERE songid=$1;`, songid)
	err := row.Scan(&name)
	return err == nil
}

// SongExtension returns song file's extension if said song exists, otherwise it
// returns a non-nil error.
func (dbi *DBI) SongExtension(songid int) (extension string, err error) {
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

// AddGenre simply adds a genre to album with given albumid. Returns error if
// something went wrong.
func (dbi *DBI) AddGenre(genre string, albumid int) (err error) {
	_, err = dbi.Exec(`INSERT INTO genres (genre, albumid) VALUES ($1, $2)`,
		genre, albumid)
	return
}

// AddFeat simply adds a feat to song with given songid. Returns error if
// something went wrong.
func (dbi *DBI) AddFeat(songid int, featid int) (err error) {
	_, err = dbi.Exec(`INSERT INTO feats (songid, featid) VALUES ($1, $2)`,
		songid, featid)
	return
}

func (dbi *DBI) albumID(songid int) (albumid int, err error) {
	row := dbi.QueryRow(`SELECT albumid FROM songs WHERE songid=$1;`, songid)
	err = row.Scan(&albumid)
	return
}
