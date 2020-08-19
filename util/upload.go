package util

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/Lisn-Rocks/server/config"
)

// AuthUpload uses Hash function to validate upload password.
func AuthUpload(r *http.Request) bool {
	password := r.FormValue("password")

	// I know that hardcoding a hash is a crude way to check
	// identity but it has to stay until we can properly handle user
	// permissions.
	if Hash(password, config.Salt) != config.Hash {
		return false
	}

	return true
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandName returns random string of letters that we use as filename generator.
func RandName(length uint) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// AlbumMeta contains album metadata.
type AlbumMeta struct {
	Album    string   `json:"album"`
	Artist   string   `json:"artist"`
	CoverExt string   `json:"coverext"`
	Genres   []string `json:"genres"`
	Songs    []struct {
		Feat     []string `json:"feat"`
		Song     string   `json:"song"`
		AudioExt string   `json:"audioext"`
	} `json:"songs"`
}

// ReadAlbumMeta returns album metadata read from "meta.json" file. It expects
// a string path to the folder that contains album files.
func ReadAlbumMeta(apath string) (data *AlbumMeta, err error) {
	metaPath := path.Join(apath, "meta.json")

	if PathNotExists(metaPath) {
		err = errors.New("did not find meta.json file")
		return
	}

	m, _ := os.Open(metaPath)
	defer m.Close()

	meta, _ := ioutil.ReadAll(m)
	data = new(AlbumMeta)
	err = json.Unmarshal(meta, data)
	return
}

// Fix is a clever function that allows admins to save time by ommitting common
// defaults like '.mp3' audioext in the 'meta.json' files.
func (meta *AlbumMeta) Fix() {
	if meta.CoverExt == "" {
		meta.CoverExt = ".jpg"
	}

	for i, song := range meta.Songs {
		if song.AudioExt == "" {
			meta.Songs[i].AudioExt = ".mp3"
		}
	}
}

// Unzip will decompress a zip archive, moving all files and folders
// within the r zip.ReadCloser to an output directory dist.
func Unzip(r *zip.ReadCloser, dest string) (filenames []string, err error) {
	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			err = fmt.Errorf("%s: illegal file path", fpath)
			return
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return
		}

		var outFile *os.File
		outFile, err = os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return
		}

		var rc io.ReadCloser
		rc, err = f.Open()
		if err != nil {
			return
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return
		}
	}

	return
}

// PathNotExists returns a boolean flag.
func PathNotExists(apath string) bool {
	_, err := os.Stat(apath)
	return os.IsNotExist(err)
}

// SaveSongs moves songs from newly extracted album into their proper location
// within the config.SongsFolder.
func SaveSongs(
	firstSongID int,
	meta *AlbumMeta,
	folderPath string,
) (err error) {
	for id, song := range meta.Songs {
		songFileName := path.Join(folderPath, song.Song+song.AudioExt)
		songidString := fmt.Sprintf("%d%s", firstSongID+id, song.AudioExt)
		songTargetName := path.Join(config.SongsFolder, songidString)

		err = os.Rename(songFileName, songTargetName)
	}

	return
}

// SaveAlbumCover moves album cover from newly extracted album into its proper
// location within the config.AlbumsFolder.
func SaveAlbumCover(albumid int, meta *AlbumMeta, folderPath string) error {
	coverFileName := path.Join(folderPath, "cover"+meta.CoverExt)
	albumidString := fmt.Sprintf("%d%s", albumid, meta.CoverExt)
	coverTargetName := path.Join(config.AlbumsFolder, albumidString)

	err := os.Rename(coverFileName, coverTargetName)
	if err != nil {
		return err
	}

	// Minimize album cover.
	minFileName := fmt.Sprintf("%d-min%s", albumid, meta.CoverExt)
	minFullPath := path.Join(config.AlbumsFolder, minFileName)
	cmd := exec.Command("convert", coverTargetName, "-resize", "100x100",
		minFullPath)
	_, err = cmd.Output()
	return err
}
