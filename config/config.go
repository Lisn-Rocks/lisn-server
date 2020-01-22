package config

import "path"


// These are project-level constants used by various packages. They serve to
// improve code scalability and ease of deployment.
const (
	// InitRequired specifies whether the song database requeres initialization.
	InitRequired = false

	// Port specifies server port.
	Port = ":8000"

	// RootFolder contains path to the root of this project. It is extremely
	// useful for the app called static that requires information about absolute
	// paths of different files on the system.
	//
	// Make sure to change it before running the server on your machine!
	RootFolder = "/home/vr4n18/go/src/github.com/sharpvik/lisn-backend"
)


// These are declared as var because their values are computed based on the
// constants above.
var (
	// TemplatesFolder contains path to the templates folder calculated based on
	// the RootFolder path.
	TemplatesFolder = path.Join(RootFolder, "templates")

	// AppsFolder contains all apps used in this project. Apps are basically
	// subroutines and handlers that provided different services, reacting to
	// users' requests.
	AppsFolder = path.Join(RootFolder, "apps")
)
