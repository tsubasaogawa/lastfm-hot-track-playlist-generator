package ytmusic

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

type Playlist struct {
	service       *youtube.Service
	Title         string
	Description   string
	PrivacyStatus string // private, public, unlisted
	Id            string
}

func NewPlaylist(svc *youtube.Service, tt string, desc string, stat string) (*Playlist, error) {
	plExists, err := exists(svc, tt)
	if plExists {
		return nil, fmt.Errorf("Playlist `%s` is already exists", tt)
	} else if err != nil {
		return nil, err
	}

	insert := svc.Playlists.Insert([]string{"snippet", "status"}, &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:       tt,
			Description: desc,
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: stat,
		},
	})

	resp, err := insert.Do()
	if err != nil {
		return nil, err
	}

	return &Playlist{
		service:       svc,
		Title:         tt,
		Description:   desc,
		PrivacyStatus: stat,
		Id:            resp.Id,
	}, nil
}

func exists(svc *youtube.Service, tt string) (bool, error) {
	list := svc.Playlists.List([]string{"snippet"}).Mine(true)
	resp, err := list.Do()
	if err != nil {
		return false, err
	}

	for _, myPlaylist := range resp.Items {
		if myPlaylist.Snippet.Title == tt {
			return true, nil
		}
	}
	return false, nil
}

func (p *Playlist) AddItem(tr *Track) error {
	insert := p.service.PlaylistItems.Insert([]string{"snippet"}, &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: p.Id,
			ResourceId: &youtube.ResourceId{
				VideoId: tr.Id,
				Kind:    "youtube#video",
			},
		},
		//
	})

	_, err := insert.Do()
	if err != nil {
		return err
	}

	return nil
}
