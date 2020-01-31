package config

import (
    "path"
    "os"
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

    // StoreFolder contains all songs that we stream.
    StoreFolder = path.Join(RootFolder, "store")

    // StaticFolder contains static files that are to be served.
    StaticFolder = path.Join(RootFolder, "static")
)


// InitRequired tell us whether the song database requeres initialization.
// To figure that out, it checks whether path DatabaseFile actually exists.
// Therefore, InitRequired must be run before
//
//     sql.Open("sqlite3", "./songs.db")
//
// is run. Otherwise, it gives falty results.
func InitRequired() bool {
    _, err := os.Stat(DatabaseFile)
    return os.IsNotExist(err)
}
