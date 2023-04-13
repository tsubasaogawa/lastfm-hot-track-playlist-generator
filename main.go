package main

import (
	"fmt"
	"log"

	"github.com/tsubasaogawa/lastfm-hot-track-playlist-generator/lastfm"
)

func main() {
	chart := lastfm.WeeklyTrackChart{}
	tracks, err := chart.GetTracks("hurikake")
	if err != nil {
		log.Fatalln(err)
	}
	for _, track := range tracks {
		fmt.Printf("%s - %s\n", track.ArtistName, track.Name)
	}
}
