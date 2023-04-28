package ytmusic

import (
	"strings"

	"google.golang.org/api/youtube/v3"
)

type Search struct {
	service         *youtube.Service
	MaxResults      int64
	MaxSearchVideos int
	searchCount     int
	Q               string
	RegionCode      string
	VideoCategoryId string
}

func NewSearch(svc *youtube.Service) *Search {
	return &Search{
		service:         svc,
		MaxResults:      5,
		MaxSearchVideos: 5,
		searchCount:     0,
		Q:               "",
		RegionCode:      "JP",
		VideoCategoryId: "10", // "Music"
	}
}

func (s *Search) Do() (*Track, error) {
	tok := ""
	mostRelatedVideo := Track{}
	for i := 1; ; i++ {
		items, nextTok, err := s.listVideos(tok)
		if err != nil {
			return nil, err
		}
		art := s.chooseArtTrack(items)
		if art != nil {
			return art, nil
		} else if i == 1 {
			mostRelatedVideo = Track{
				Title:  items[0].Snippet.Title,
				Artist: items[0].Snippet.ChannelTitle,
				Id:     items[0].Id.VideoId,
			}
		}

		tok = nextTok
		if tok == "" {
			break
		}
	}

	// when search results are non-art tracks only
	return &mostRelatedVideo, nil
}

func (s *Search) listVideos(tok string) ([]*youtube.SearchResult, string, error) {
	print("nextTok = " + tok + "\n")
	search := s.service.Search.List([]string{"snippet"}).
		MaxResults(s.MaxResults).
		Q(s.Q).
		Type("video").
		VideoCategoryId(s.VideoCategoryId).
		RegionCode(s.RegionCode).
		PageToken(tok)

	resp, err := search.Do()
	if err != nil {
		return nil, "", err
	}

	return resp.Items, resp.NextPageToken, nil
}

func (s *Search) chooseArtTrack(items []*youtube.SearchResult) *Track {
	for _, item := range items {
		s.searchCount++
		print("Try: " + item.Snippet.Title + "\n")
		if !s.isArtTrack(item.Snippet) {
			if s.searchCount <= s.MaxSearchVideos {
				continue
			} else {
				break
			}
		}
		return &Track{
			Title:  item.Snippet.Title,
			Artist: item.Snippet.ChannelTitle,
			Id:     item.Id.VideoId,
		}
	}
	return nil
}

func (s *Search) isArtTrack(snip *youtube.SearchResultSnippet) bool {
	return strings.HasSuffix(snip.ChannelTitle, "- Topic") ||
		strings.HasPrefix(snip.Description, "Provided to YouTube")
}
