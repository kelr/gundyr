// An example to obtain a user authentication token for user's email.
// Uses the token to get info about the user.
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kelr/go-twitch-api/helix"
	"golang.org/x/oauth2"
	"time"
)

// Provide your Client ID and secret. Set your redirect URI to one that you own.
// Better to set these as environment variables.
const (
	clientID     = ""
	clientSecret = ""
	redirectURI  = ""
)

// Set scopes to request from the user
var scopes = []string{"user:read:broadcast"}

func main() {
	// Setup OAuth2 config
	config, err := helix.NewUserAuth(clientID, clientSecret, redirectURI, &scopes)
	if err != nil {
		fmt.Println(err)
	}

	// Import an existing token to use
	token := new(oauth2.Token)
	token.AccessToken = ""
	token.Expiry = time.Date(2020, 5, 14, 6, 45, 0, 0, time.UTC)
	token.RefreshToken = ""
	token.TokenType = "bearer"

	// Create API client. User token will be automatically refreshed.
	client, err := helix.NewTwitchClientUserAuth(config, token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get a list of all active extensions for the user matching the token
	resp, err := client.GetUserActiveExtensions(nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print
	obj, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(obj))
}
