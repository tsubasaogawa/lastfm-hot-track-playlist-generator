package ytmusic

import (
	"strings"

	"google.golang.org/api/youtube/v3"
)

type Search struct {
	service         *youtube.Service
	MaxResults      int64
	Repeats         int
	Q               string
	RegionCode      string
	VideoCategoryId string
}

func NewSearch(svc *youtube.Service) *Search {
	return &Search{
		service:         svc,
		MaxResults:      5,
		Repeats:         3,
		Q:               "",
		RegionCode:      "JP",
		VideoCategoryId: "10", // "Music"
	}
}

func (s *Search) Do() (*Track, error) {
	nextTok := ""
	mostRelatedTrack := Track{}
	for i := 0; ; i++ {
		search := s.service.Search.List([]string{"snippet"}).
			MaxResults(s.MaxResults).
			Q(s.Q).
			// FIXME: Error `googleapi: Error 400: Request contains an invalid argument., badRequest` occurs. For me only?
			// VideoCategoryId(s.VideoCategoryId).
			RegionCode(s.RegionCode).
			PageToken(nextTok)

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

		if i == 0 {
			mostRelatedTrack = Track{
				Title:  resp.Items[0].Snippet.Title,
				Artist: resp.Items[0].Snippet.ChannelTitle,
				Id:     resp.Items[0].Id.VideoId,
			}
		}

		nextTok = resp.NextPageToken
		if nextTok == "" {
			break
		}
	}

	// when search results are non-art tracks only
	return &mostRelatedTrack, nil
}

func (s *Search) isArtTrack(snip *youtube.SearchResultSnippet) bool {
	return strings.HasSuffix(snip.ChannelTitle, "- Topic") || strings.HasPrefix(snip.Description, "Provided to YouTube")
}
