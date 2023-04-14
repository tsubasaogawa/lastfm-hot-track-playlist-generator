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
	for _, track := range tracks {
		fmt.Printf("%s - %s\n", track.ArtistName, track.Name)
	}

	service, err := ytmusic.NewService(os.Getenv("YT_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}
	search := ytmusic.NewSearch(service)
	search.Q = "Pixies"
	searchItems := search.Do()
	for _, item := range searchItems {
		fmt.Printf("%s - %s\n", item.Title, item.Id)
	}
}
