# gundyr

Gundyr provides an easy to use interface to the [Helix Twitch API](https://dev.twitch.tv/docs/api/reference) and [Twitch PubSub](https://dev.twitch.tv/docs/pubsub).

It handles both app access as well as user access tokens. All tokens used are automatically refreshed.

Note: This is a work in progress and a project to help me learn Go :). May not provide full functionality.

## Install

```bash
$ go get github.com/kelr/gundyr
```

## Usage
Example to convert a Twitch login name to account ID

```go
package main

import (
	"log"
	"github.com/kelr/gundyr"
)

// Provide your Client ID and secret here.
const (
	clientID     = ""
	clientSecret = ""
)

func main() {
	c, err := gundyr.NewHelix(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	userID, err := c.UserToID("kyrotobi")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(userID)
}
```

## Contributions
Any and all contributions or bug fixes are appreciated.
