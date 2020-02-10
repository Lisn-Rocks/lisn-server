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
    // DatabaseFile contains path to the main songs databse.
    DatabaseFile = path.Join(RootFolder, "songs.db")

    // TemplatesFolder contains path to the templates folder calculated based on
    // the RootFolder path.
    TemplatesFolder = path.Join(RootFolder, "templates")

    // AppsFolder contains all apps used in this project. Apps are basically
    // subroutines and handlers that provided different services, reacting to
    // users' requests.
    AppsFolder = path.Join(RootFolder, "apps")

	// SongsFolder contains all data about the songs we stream.
	SongsFolder = path.Join(RootFolder, "songs")
	
	// SongsAudioFolder contains all the audio files of the songs we serve.
	SongsAudioFolder = path.Join(SongsFolder, "audio")

	// SongsCoverFolder contains all album / playlist covers for the songs.
	SongsCoverFolder = path.Join(SongsFolder, "cover")

    // PublicFolder contains static files that are to be served publically.
    // These files do not contain any sensitive data and thus we don't really
    // care if they can be accessed arbitrarily without any identity checkup.
    PublicFolder = path.Join(RootFolder, "public")
)
