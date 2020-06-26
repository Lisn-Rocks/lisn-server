package dbi

import (
	"os"

	"github.com/sharpvik/lisn-server/util"
)

// UploadAlbum function deals with album uploads in terms of database operations
// and storage manipulations.
func (dbi *DBI) UploadAlbum(archive *os.File, meta *util.AlbumMeta) (re error) {
	artistid, re := dbi.SearchAddArtist(meta.Artist)
	if re != nil {
		return
	}

	albumid, re := dbi.SearchAddAlbum(meta.Album, artistid, meta.CoverExt)
	if re != nil {
		return
	}

	for _, genre := range meta.Genres {
		err := dbi.AddGenre(genre, albumid)
		if err != nil {
			dbi.logr.Printf("failed to add genre: %s", genre)
		}
	}

	for _, song := range meta.Songs {
		songid, err := dbi.AddSong(song.Song, albumid, song.AudioExt)
		if err != nil {
			dbi.logr.Printf("failed to add song: %s", song.Song)
			continue
		}

		for _, feat := range song.Feat {
			featid, err := dbi.SearchAddArtist(feat)
			if err != nil {
				dbi.logr.Printf("failed to add feat artist: %s", feat)
				continue
			}
			err = dbi.AddFeat(songid, featid)
			if err != nil {
				dbi.logr.Printf("failed to add feat: %s", feat)
				continue
			}
		}
	}

	return
}
