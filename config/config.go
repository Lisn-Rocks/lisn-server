package config

import (
    "path"
    "os"
)


// These are project-level constants used by various packages. They serve to
// improve code scalability and ease of deployment.
const (
    // Host specifies server host address.
    Host = "localhost"

    // Port specifies server port.
    Port = ":8000"

    // RootFolder contains path to the root of this project. It is extremely
    // useful for the app called static that requires information about absolute
    // paths of different files on the system.
    //
    // MAKE SURE TO CHANGE IT BEFORE RUNNING THE SERVER ON YOUR MACHINE!
    RootFolder = "/home/sharpvik/Public/Lisn"
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

    // PublicFolder contains static files that are to be served publically.
    // These files do not contain any sensitive data and thus we don't really
    // care if they can be accessed arbitrarily without any identity checkup.
    PublicFolder = path.Join(RootFolder, "pub")

    // FailFolder contains static HTML files that are sent to user whenever some
    // sort of failure occurs. It has files named '<ERROR_CODE>.html' for every
    // error code utilized by the app.
    FailFolder = path.Join(PublicFolder, "fail")

    // LisnFolder contains static files that, in their entirity, make up the
    // whole of the Lisn Music Streaming App.
    LisnFolder = path.Join(PublicFolder, "lisn")
)


// The following values are used by the logger. Configure them as you please.

// LogFile contains path to the log file
var LogFile = path.Join(RootFolder, "log")

// LogWriter must implement Writer and is uded by the logger to know where to
// write the logs to. Uncomment the one you wish to use.
var LogWriter = os.Stdout
// var LogWriter, _ = os.Create(LogFile)

// LogPrefix is a string used by the logger to prefix every log message.
const LogPrefix = ""
