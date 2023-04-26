package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Artist struct {
	ArtistName string `json:"text"`
}

type Track struct {
	Artist    `json:"artist,string"`
	Name      string `json:"name"`
	Playcount int    `json:"playcount,string"`
}

type _WeeklyTrackChart struct {
	Tracks []Track `json:"track,string"`
}

type WeeklyTrackChart struct {
	_WeeklyTrackChart `json:"weeklytrackchart,string"`
}

func GetTracks(user string, apikey string, from int64, to int64) ([]Track, error) {
	if apikey == "" {
		return nil, fmt.Errorf("Last.FM API Key is required")
	}
	ENDPOINT := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=user.getweeklytrackchart&user=%s&api_key=%s&format=json&from=%d&to=%d&limit=5", user, apikey, from, to)

	req, err := http.NewRequest(http.MethodGet, ENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	invalidByteJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	validJson := strings.ReplaceAll(string(invalidByteJson), "#text", "text")

	weekly := WeeklyTrackChart{}
	err = json.Unmarshal([]byte(validJson), &weekly)

	return weekly.Tracks, err
}
