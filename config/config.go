package config

import (
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
    RootFolder = "/home/sharpvik/go/src/github.com/sharpvik/Lisn"
)


// These are declared as var because their values are computed based on the
// constants above.
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
    PublicFolder = path.Join(RootFolder, "public")
)
