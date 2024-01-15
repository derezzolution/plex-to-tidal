package tidal

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func (a *AuthResponse) IsAuthenticated() bool {
	return len(a.AccessToken) > 0
}

type Resource struct {
	ArtifactType  string                 `json:"artifactType"`
	ID            string                 `json:"id"`
	Title         string                 `json:"title"`
	Artists       []Artist               `json:"artists"`
	Album         Album                  `json:"album"`
	Duration      int                    `json:"duration"`
	TrackNumber   int                    `json:"trackNumber"`
	VolumeNumber  int                    `json:"volumeNumber"`
	Isrc          string                 `json:"isrc"` // Potential primary key
	Copyright     string                 `json:"copyright"`
	MediaMetadata Metadata               `json:"mediaMetadata"`
	Properties    map[string]interface{} `json:"properties"`
	TidalURL      string                 `json:"tidalUrl"`
}

type Artist struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Picture []string `json:"picture"`
	Main    bool     `json:"main"`
}

type Album struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	ImageCover []Image  `json:"imageCover"`
	VideoCover []string `json:"videoCover"`
}

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Metadata struct {
	Tags []string `json:"tags"`
}

type Track struct {
	Resource Resource `json:"resource"`
}

type Playlist struct {
	// TODO: Populate after we get a known-working response
}
