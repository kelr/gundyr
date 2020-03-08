package twitchapi

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
	"fmt"
	"context"
)

const (
	helixRootURL = "https://api.twitch.tv/helix"
)

// Handles communication with the Twitch API.
type TwitchClient struct {
	conn *http.Client
	ClientID string
	ClientSecret string
}

// Returns a new Twitch Client. If clientID is "", it will not be appended on the request header.
// A client credentials config is established which auto-refreshes OAuth2 access tokens
// Currently ONLY uses Client Credentials flow. Not intended for user access tokens.
func NewTwitchClient(clientID string, clientSecret string) *TwitchClient {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		fmt.Println("Error in getting a token: ", err)
	}

	return &TwitchClient{
		ClientID: clientID,
		ClientSecret: clientSecret,
		conn: config.Client(context.Background()),
	}
}


// Create and send an HTTP request.
func (client *TwitchClient) sendRequest(path string, params interface{}, result interface{}) (*http.Response, error) {
	targetUrl, err := url.Parse(helixRootURL + path)
	if err != nil {
		return nil, err
	}

	// Convert optional params to URL queries
	if params != nil {
		qs, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		targetUrl.RawQuery = qs.Encode()
	}

	request, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add on optional headers
	if client.ClientID != "" {
		request.Header.Set("Client-ID", client.ClientID)
	}

	// Send the request
	resp, err := client.conn.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(result)
	if err == io.EOF {
		err = nil
	}

	// TODO: Check response code
	return resp, nil
}
