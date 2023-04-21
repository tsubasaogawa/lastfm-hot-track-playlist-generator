package ytmusic

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

func NewService(oauthfile string) (*youtube.Service, error) {
	oauthjson, err := os.ReadFile(oauthfile)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(oauthjson, youtube.YoutubeReadonlyScope)
	if err != nil {
		return nil, err
	}

	url := config.AuthCodeURL("test", oauth2.AccessTypeOffline)
	fmt.Println(url)

	var s string
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	}

	oauthtoken, err := config.Exchange(oauth2.NoContext, s)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client := config.Client(ctx, oauthtoken)
	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}

	return service, nil
}
