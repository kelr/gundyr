// An example to obtain a user authentication token for user's email.
// Uses OAuth2 Authorization Code Flow
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kelr/go-twitch-api/helix"
)

// Provide your Client ID and secret. Set your redirect URI to one that you own.
// The URI must match exactly with the one registered by your app on the Twitch Developers site
const (
	clientID       = ""
	clientSecret   = ""
	redirectURI    = "http://localhost"
	targetUsername = ""
)

// Set scopes to request from the user
var scopes = []string{"user:read:email"}

func main() {
	// Setup OAuth2 configuration
	config, err := helix.NewUserAuth(clientID, clientSecret, redirectURI, &scopes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the URL to send to the user and the state code to protect against CSRF attacks.
	url, state := helix.GetAuthCodeURL(config)
	fmt.Println(url)
	fmt.Println("Ensure that state recieved at URI is:", state)

	// Enter the code received by the redirect URI. Ensure that the state value
	// obtained at the redirect URI matches the previous state value.
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Println(err)
	}

	// Obtain the user token through the code. This token can be reused as long as
	// it has not expired, but the auth code cannot be reused.
	token, err := helix.TokenExchange(config, authCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the API client. User token will be automatically refreshed.
	client, err := helix.NewTwitchClientUserAuth(config, token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get user information, will include email for the user you have a token from
	opt := &helix.GetUsersOpt{
		Login: targetUsername,
	}

	response, err := client.GetUsers(opt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print
	obj, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(obj))
}
