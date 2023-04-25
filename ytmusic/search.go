package ytmusic

import (
	"strings"

	"google.golang.org/api/youtube/v3"
)

type Search struct {
	service         *youtube.Service
	MaxResults      int64
	Q               string
	RegionCode      string
	VideoCategoryId string
}

func NewSearch(svc *youtube.Service) *Search {
	return &Search{
		service:         svc,
		MaxResults:      5,
		Q:               "",
		RegionCode:      "JP",
		VideoCategoryId: "10", // "Music"
	}
}

func (s *Search) Do() (*Track, error) {
	search := s.service.Search.List([]string{"snippet"}).
		MaxResults(s.MaxResults).
		Q(s.Q).
		// FIXME: Error `googleapi: Error 400: Request contains an invalid argument., badRequest` occurs. For me only?
		// VideoCategoryId(s.VideoCategoryId).
		RegionCode(s.RegionCode)

	resp, err := search.Do()
	if err != nil {
		return nil, err
	}

	for _, item := range resp.Items {
		if !s.isArtTrack(item.Snippet) {
			continue
		}
		return &Track{
			Title:  item.Snippet.Title,
			Artist: item.Snippet.ChannelTitle,
			Id:     item.Id.VideoId,
		}, nil
	}

	return &Track{
		Title:  resp.Items[0].Snippet.Title,
		Artist: resp.Items[0].Snippet.ChannelTitle,
		Id:     resp.Items[0].Id.VideoId,
	}, nil
}

func (s *Search) isArtTrack(snip *youtube.SearchResultSnippet) bool {
	return strings.HasSuffix(snip.ChannelTitle, "- Topic") || strings.HasPrefix(snip.Description, "Provided to YouTube")
}
