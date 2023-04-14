package ytmusic

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func NewService(apikey string) (*youtube.Service, error) {
	ctx := context.Background()
	if apikey == "" {
		return nil, fmt.Errorf("YouTube API Key is required.")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(apikey))
	if err != nil {
		return nil, err
	}

	return service, nil
}
