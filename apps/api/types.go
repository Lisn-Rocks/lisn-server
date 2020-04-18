package api


type song struct {
    SongID  int    `json:"songid"`
    Name    string `json:"name"`
    Artist  string `json:"artist"`
    Genre   string `json:"genre"`
    Album   string `json:"album"`
}
