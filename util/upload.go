package util

import (
	"archive/zip"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
)

// Hash returns base64 encoded SHA512 salted hash of password as a string.
func Hash(password, salt string) string {
	h := sha512.Sum512([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(h[:])
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandName returns random string of letters that we use as filename generator.
func RandName() string {
	b := make([]byte, 32)
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

// Meta contains album metadata.
type Meta struct {
	Album    string   `json:"album"`
	Artist   string   `json:"artist"`
	Coverext string   `json:"coverext"`
	Genres   []string `json:"genres"`
	Songs    []struct {
		Song    string `json:"song"`
		Songext string `json:"songext"`
	} `json:"songs"`
}

// ReadMeta returns album metadata read from "meta.json" file.
func ReadMeta(apath string) (data *Meta, err error) {
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

	data = new(Meta)
	err = json.Unmarshal(meta, data)
	return
}
