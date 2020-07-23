package router

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/sharpvik/lisn-server/config"
	"github.com/sharpvik/lisn-server/util"
)

func processUpload(w http.ResponseWriter, r *http.Request, e *Env) (re error) {
	e.logr.Print("incoming upload request")

	// Parse form and check authenticity.
	if err := r.ParseMultipartForm(config.MaxMemUploadSize); err != nil {
		e.logr.Println(err)
	}

	if !util.AuthUpload(r) {
		e.logr.Print("authentication failed (hash did not match)")
		fmt.Fprint(w, "<h1>You aren't authorized to upload.</h1>")
		return
	}

	album, _, err := r.FormFile("album")
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to retreive archive.</h1>")
		return
	}
	defer album.Close()

	// Save archive under random filename for further reading.
	randName := util.RandName(32)
	folderName := path.Join(config.ArchivesFolder, randName)
	os.Mkdir(folderName, os.ModeDir|os.ModePerm)
	apath := path.Join(config.ArchivesFolder, randName+".zip")

	archive, err := os.Create(apath)
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to create archive file.</h1>")
		return
	}
	defer os.Remove(apath)

	io.Copy(archive, album)
	archive.Close()

	rdr, err := zip.OpenReader(apath)
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to read archive.</h1>")
		return
	}

	// Unzip archive into folder named randName.
	util.Unzip(rdr, folderName)
	defer os.RemoveAll(randName)

	meta, err := util.ReadAlbumMeta(folderName)
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to read metadata.</h1>")
		return
	}

	firstSongID, albumid, err := e.dbi.UploadAlbum(meta)
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to process upload.</h1>")
		return
	}

	err = util.SaveSongs(firstSongID, meta, folderName)
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to save songs.</h1>")
	}

	err = util.SaveAlbumCover(albumid, meta, folderName)
	if err != nil {
		e.logr.Println(err)
		fmt.Fprint(w, "<h1>Failed to save album cover.</h1>")
	}

	fmt.Fprint(w, "<h1>Your contribution is greatly appreciated!</h1>")
	return
}
