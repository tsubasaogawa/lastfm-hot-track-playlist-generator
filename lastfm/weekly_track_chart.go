package lastfm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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

var (
	LASTFM_API_KEY   string = os.Getenv("LASTFM_API_KEY")
	LASTFM_USER_NAME string = os.Getenv("LASTFM_USER_NAME")
)

func (chart *WeeklyTrackChart) GetTracks(user string) ([]Track, error) {
	ENDPOINT := "http://ws.audioscrobbler.com/2.0/?method=user.getweeklytrackchart&user=" + user + "&api_key=" + LASTFM_API_KEY + "&format=json&from=1648738800&to=1680274800&limit=5"

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
	if err != nil {
		return nil, err
	}

	return weekly.Tracks, nil
}
