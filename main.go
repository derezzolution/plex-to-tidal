package main

import (
	"embed"
	"log"

	"github.com/derezzolution/go-plex-client"

	"github.com/derezzolution/plex-to-tidal/config"
	"github.com/derezzolution/plex-to-tidal/tidal"
)

//go:embed config.json
var packageFS embed.FS

func main() {
	config, err := config.NewConfig(&packageFS)
	if err != nil {
		log.Fatal("error reading config:", err)
	}

	plexConnection, err := plex.New(config.PlexServerUrl, config.PlexToken)
	if err != nil {
		log.Fatal("error connecting to plex server:", err)
	}

	playlist, err := plexConnection.GetPlaylist(config.PlexPlaylistKey)
	if err != nil {
		log.Fatal("error fetching playlist:", err)
	}
	log.Printf(`loaded playlist "%s", found %d tracks`, playlist.MediaContainer.Title,
		len(playlist.MediaContainer.Metadata))

	// Debugging: Plex track dump
	i := 0
	for _, item := range playlist.MediaContainer.Metadata {
		if i == 0 {
			log.Println("item ==> ", item)
			i++
		} else {
			log.Printf("rating key %s", item.RatingKey)
		}
	}

	// Debugging: Tidal track dump
	t := &tidal.TidalClient{}
	t.Authenticate(config)
	track, err := t.LookupTrackByID("221266184", "US")
	if err != nil {
		log.Fatalf("unable to lookup track: %s", err)
	}
	log.Println("track => ", track)
}
