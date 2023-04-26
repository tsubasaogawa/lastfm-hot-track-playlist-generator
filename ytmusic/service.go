package ytmusic

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const (
	TOKEN_FILE = "token.json"
)

func NewService(oauthfile string) (*youtube.Service, error) {
	oauthjson, err := os.ReadFile(oauthfile)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(oauthjson, youtube.YoutubeScope)
	if err != nil {
		return nil, err
	}

	token, err := readSavedToken()
	if err != nil {
		token, err = generateToken(config)
		if err != nil {
			return nil, err
		}
		log.Println("Generated token is saved to " + TOKEN_FILE)
	}

	ctx := context.Background()
	client := config.Client(ctx, token)
	service, err := youtube.New(client)

	return service, err
}

func readSavedToken() (*oauth2.Token, error) {
	f, err := os.Open(TOKEN_FILE)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	token := oauth2.Token{}
	err = json.NewDecoder(f).Decode(&token)

	return &token, err
}

func generateToken(c *oauth2.Config) (*oauth2.Token, error) {
	url := c.AuthCodeURL("test", oauth2.AccessTypeOffline) // FIXME
	fmt.Printf("Access to the following URL: \n%s\nAuth code is: ", url)

	var s string
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	}

	token, err := c.Exchange(oauth2.NoContext, s)
	if err != nil {
		return nil, err
	}

	// save
	f, err := os.OpenFile(TOKEN_FILE, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)

	return token, err
}
