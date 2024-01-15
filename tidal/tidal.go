package tidal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/derezzolution/plex-to-tidal/config"
)

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
