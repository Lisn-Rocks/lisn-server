package songinfo


type song struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Genre    string `json:"genre"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
}