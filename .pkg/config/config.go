// Package config provides functions to retreive appropriate info from env variables,
// structs to store env variables values and constants that will be used during execution
package config

import (
	"os"
	"path"
)

// These are project-level constants used by various packages. They serve to
// improve code scalability and ease of deployment.
const (
	// Port specifies server port.
	Port = ":8000"

	// RootFolder contains path to the root of this project. It is extremely
	// useful for the app called static that requires information about absolute
	// paths of different files on the system.
	//
	// MAKE SURE TO CHANGE IT BEFORE RUNNING THE SERVER ON YOUR MACHINE!
	RootFolder = "/home/sharpvik/Public/lisn"
)

// The following variables contain string paths of different folders.
// These are declared as var because their values are computed at runtime based
// on the RootFolder value.
var (
	// StorageFolder contains all data necessary for the service to run properly.
	StorageFolder = path.Join(RootFolder, "storage")

	// SongsFolder contains all the audio files of the songs we serve.
	SongsFolder = path.Join(StorageFolder, "songs")

	// AlbumsFolder contains all album / playlist covers for the songs.
	AlbumsFolder = path.Join(StorageFolder, "albums")

	// ArchivesFolder contains temporary album archives.
	ArchivesFolder = path.Join(StorageFolder, "archives")

	// PublicFolder contains static files that are to be served publically.
	// These files do not contain any sensitive data and thus we don't really
	// care if they can be accessed arbitrarily without any identity checkup.
	PublicFolder = path.Join(RootFolder, "pub")

	// FailFolder contains static HTML files that are sent to user whenever some
	// sort of failure occurs. It has files named '<ERROR_CODE>.html' for every
	// error code utilized by the app.
	FailFolder = path.Join(PublicFolder, "fail")

	// UploadFolder contains static HTML files that are sent to user on request
	// to upload a new album.
	UploadFolder = path.Join(PublicFolder, "upload")

	// LisnFolder contains static files that, in their entirity, make up the
	// whole of the Lisn Music Streaming App.
	LisnFolder = path.Join(PublicFolder, "lisn")

	// LogsFolder contains all the log files related to Lisn Music Streaming
	// Service activities (including, for example, CD build logs).
	LogsFolder = path.Join(RootFolder, "logs")
)

// The following values are used by the logger. Configure them as you please.

var (
	// LogFile contains path to the log file
	LogWriter, _ = os.Create(LogFile)

	// LogWriter must implement Writer and is used by the logger to know where to
	// write the logs to.
	LogFile = path.Join(LogsFolder, "lisn-server.log")
)

const (
	// LogPrefix is a string used by the logger to prefix every log message.
	LogPrefix = ""
	// MaxMemUploadSize is the number of bytes that will be saved to memory from
	// multipart upload form.
	MaxMemUploadSize int64 = 500000000
)

// Everything related to db config

// DB contains all the info required to connect to Postgre database
type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// Lisn stores a struct containing all the info about database
var Lisn DB

// Hash functions

// TODO Possibly add these as env var as well
const (
	// Hash is a salted hash of the master password that allows to upload albums
	// to server using the "/upload" site.
	Hash = `averylongstringofsymbolsthatrepresentsavalidsha512hashdigestedthroughbase64encodingmustgohere`

	// Salt is a salt appended to master password before hashing.
	Salt = `anotherlongsequenceofrandomsymbolspreferrably32orlonger`
)
