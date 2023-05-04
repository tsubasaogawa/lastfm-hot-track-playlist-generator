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

func NewPlaylist(svc *youtube.Service, tt string, desc string, stat string, dup bool, max int64) (*Playlist, error) {
	switch stat {
	case "private", "public", "unlisted":
	default:
		return nil, fmt.Errorf("PrivacyStatus must be private, public or unlisted")
	}

	_exists, err := exists(svc, tt, max)
	if _exists && !dup {
		return nil, fmt.Errorf("Playlist `%s` already exists", tt)
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

func exists(svc *youtube.Service, tt string, max int64) (bool, error) {
	tok := ""
	for {
		items, nextTok, err := listPlaylists(svc, tok, max)
		if err != nil {
			return false, err
		}
		for _, pl := range items {
			if pl.Snippet.Title == tt {
				return true, nil
			}
		}

		tok = nextTok
		if tok == "" {
			break
		}
	}

	return false, nil
}

func listPlaylists(svc *youtube.Service, tok string, max int64) ([]*youtube.Playlist, string, error) {
	list := svc.Playlists.List([]string{"snippet"}).
		Mine(true).
		MaxResults(max).
		PageToken(tok)

	resp, err := list.Do()
	if err != nil {
		return nil, "", err
	}

	return resp.Items, resp.NextPageToken, nil
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

	return err
}
