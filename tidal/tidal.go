package tidal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/derezzolution/plex-to-tidal/config"
)

// TIDAL API Reference: https://developer.tidal.com/apiref

type TidalClient struct {
	accessToken *AuthResponse
}

// Authenticate authenticates and obtains an access token.
func (t *TidalClient) Authenticate(config *config.Config) error {
	authURL := "https://auth.tidal.com/v1/oauth2/token"
	data := url.Values{
		"client_id":     {config.TidalClientId},
		"client_secret": {config.TidalClientSecret},
		"grant_type":    {"client_credentials"},
	}

	resp, err := http.Post(authURL, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		t.accessToken = &AuthResponse{}
		return err
	}
	defer resp.Body.Close()

	authResp := &AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		t.accessToken = &AuthResponse{}
		return err
	}

	t.accessToken = authResp
	return nil
}

// LookupTrackByID looks up a track by ID.
func (t *TidalClient) LookupTrackByID(trackID string, countryCode string) (*Track, error) {
	if !t.accessToken.IsAuthenticated() {
		return nil, fmt.Errorf("error looking up track: not authenticated")
	}

	trackURL := fmt.Sprintf("https://openapi.tidal.com/tracks/%s?countryCode=%s", trackID, countryCode)
	req, err := http.NewRequest("GET", trackURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.tidal.v1+json")
	req.Header.Set("Content-Type", "application/vnd.tidal.v1+json")
	req.Header.Set("Authorization", "Bearer "+t.accessToken.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	track := &Track{}
	err = json.NewDecoder(resp.Body).Decode(&track)
	if err != nil {
		return nil, err
	}

	return track, nil
}

// CreateNewPlaylist creates a new playlist on Tidal.
func (t *TidalClient) CreateNewPlaylist(name string) (*Playlist, error) {
	if !t.accessToken.IsAuthenticated() {
		return nil, fmt.Errorf("error creating playlist: not authenticated")
	}

	// I think the whole tidal api is read-only >:(
	createPlaylistURL := "https://openapi.tidal.com/playlists" // TODO - 404
	data := map[string]interface{}{
		"title":         name,
		"public":        false,
		"collaborative": false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", createPlaylistURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Debugging
	log.Printf("sending in -> %s", bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/vnd.tidal.v1+json")
	req.Header.Set("Authorization", "Bearer "+t.accessToken.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Debugging
	log.Println("Response Status Code:", resp.StatusCode)
	log.Println("Response Body:", string(bodyContent))

	// Decode the JSON content
	playlist := &Playlist{}
	err = json.Unmarshal(bodyContent, &playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}
