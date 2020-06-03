package util

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/sharpvik/lisn-server/config"
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

func findFile(files []*zip.File, name string) *zip.File {
	for _, f := range files {
		if f.Name == name {
			return f
		}
	}
	return nil
}

// UploadMeta contains album metadata.
type UploadMeta struct {
	Album    string   `json:"album"`
	Artist   string   `json:"artist"`
	Coverext string   `json:"coverext"`
	Genres   []string `json:"genres"`
	Songs    []struct {
		Feat    []string `json:"feat"`
		Song    string   `json:"song"`
		Songext string   `json:"songext"`
	} `json:"songs"`
}

// ReadUploadMeta returns album metadata read from "meta.json" file.
func ReadUploadMeta(apath string) (data *UploadMeta, err error) {
	r, err := zip.OpenReader(apath)
	defer r.Close()
	if err != nil {
		return
	}

	metaptr := findFile(r.File, "meta.json")
	if metaptr == nil {
		err = errors.New("no metadata in album")
		return
	}
	m, _ := metaptr.Open()
	meta, _ := ioutil.ReadAll(m)

	data = new(UploadMeta)
	err = json.Unmarshal(meta, data)
	return
}
