package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tsubasaogawa/lastfm-hot-track-playlist-generator/lastfm"
	"github.com/tsubasaogawa/lastfm-hot-track-playlist-generator/ytmusic"
)

func main() {
	tracks, err := lastfm.GetTracks("hurikake", os.Getenv("LASTFM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	service, err := ytmusic.NewService()
	if err != nil {
		log.Fatalln(err)
	}
	search := ytmusic.NewSearch(service)

	for _, track := range tracks {
		search.Q = fmt.Sprintf("%s - %s", track.ArtistName, track.Name)
		searchItem, err := search.Do()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s - %s (%s)\n", searchItem.Title, searchItem.Artist, searchItem.Id)
	}
}
