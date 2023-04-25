package ytmusic

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const (
	OAUTH_FILE = "oauth.json"
	TOKEN_FILE = "token.json"
)

func NewService() (*youtube.Service, error) {
	oauthjson, err := os.ReadFile(OAUTH_FILE)
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
	}

	ctx := context.Background()
	client := config.Client(ctx, token)
	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func readSavedToken() (*oauth2.Token, error) {
	f, err := os.Open(TOKEN_FILE)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	token := oauth2.Token{}
	err = json.NewDecoder(f).Decode(&token)

	return &token, nil
}

func generateToken(c *oauth2.Config) (*oauth2.Token, error) {
	url := c.AuthCodeURL("test", oauth2.AccessTypeOffline)
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
	defer f.Close()
	json.NewEncoder(f).Encode(token)

	if err != nil {
		return nil, err
	}

	return token, nil
}
