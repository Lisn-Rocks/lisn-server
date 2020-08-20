package config

// Init fills in the structs with appropriate values
func Init() {
	godotenv.Load()

	Lisn = DB{
		Host:     getOr("POSTGRES_HOST", "localhost"),
		Port:     5432,
		User:     getMust("POSTGRES_USER"),
		Password: getMust("POSTGRES_PASSWORD"),
		Name:     getMust("POSTGRES_DB"),
	}
}

// get attempts to load a value for a specific env key and returns a boolean
// whether it was successfuul or not
func get(key string) (string, bool) {
	value := os.Getenv(key)

	if len(value) == 0 {
		fmt.Printf("Failed to load: " + key)
		return "", false
	}

	return value, true
}

// getOr uses get and if no value was loaded, returns alternative value provided
func getOr(key string, alt string) string {
	if value, ok := get(key); ok {
		return value
	}

	return alt
}

// getMust uses get and if no value was loaded, panics
func getMust(key string) string {
	value, ok := get(key)
	if !ok {
		panic("Key " + key + " must be provided")
	}

	return value
}
