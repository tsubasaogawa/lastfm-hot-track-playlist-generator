package lastfm

import (
	"fmt"
)

type Artist struct {
	ArtistName string `json:"text"`
}

type Track struct {
	Artist    `json:"artist,string"`
	Name      string `json:"name"`
	Playcount int    `json:"playcount,string"`
}

func (tr *Track) Print() {
	fmt.Printf("%s - %s (%d times)", tr.Name, tr.ArtistName, tr.Playcount)
}
