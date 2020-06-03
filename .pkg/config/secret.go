package config

// Database login information.
const (
	DBhost     = "localhost"
	DBport     = 5432
	DBuser     = "user"
	DBpassword = "***"
	DBname     = "Lisn"
)

const (
	// Hash is a salted hash of the master password that allows to upload albums
	// to server using the "/upload" site.
	Hash = `averylongstringofsymbolsthatrepresentsavalidsha512hashdigestedthroughbase64encodingmustgohere`

	// Salt is a salt appended to master password before hashing.
	Salt = `anotherlongsequenceofrandomsymbolspreferrably32orlonger`
)
