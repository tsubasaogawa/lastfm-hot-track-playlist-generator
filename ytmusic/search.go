package ytmusic

import (
	"log"

	"google.golang.org/api/youtube/v3"
)

type Search struct {
	service         *youtube.Service
	MaxResults      int64
	Q               string
	RegionCode      string
	Type            string
	VideoCategoryId string
}

type SearchResultItem struct {
	Title string
	Id    string
}

func NewSearch(svc *youtube.Service) *Search {
	return &Search{
		service:         svc,
		MaxResults:      5,
		Q:               "",
		RegionCode:      "JP",
		Type:            "video",
		VideoCategoryId: "10", // "Music"
	}
}

func (s *Search) Do() []SearchResultItem {
	search := s.service.Search.List([]string{"snippet"}).
		MaxResults(s.MaxResults).
		Q(s.Q).
		RegionCode(s.RegionCode).
		Type(s.Type)

	resp, err := search.Do()
	if err != nil {
		log.Fatalln(err)
	}

	items := []SearchResultItem{}
	for _, item := range resp.Items {
		items = append(items, SearchResultItem{
			Title: item.Snippet.Title,
			Id:    item.Snippet.ChannelId,
		})
	}

	return items
}
