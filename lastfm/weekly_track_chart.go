package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type _WeeklyTrackChart struct {
	Tracks []Track `json:"track,string"`
}

type WeeklyTrackChart struct {
	_WeeklyTrackChart `json:"weeklytrackchart,string"`
}

func GetTracks(user string, apikey string, from int64, to int64, max int) ([]Track, error) {
	endpoint := fmt.Sprintf("%s?method=user.getweeklytrackchart&user=%s&api_key=%s&format=json&from=%d&to=%d&limit=%d",
		ENDPOINT_BASE, user, apikey, from, to, max)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	invalid, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	valid := strings.ReplaceAll(string(invalid), "#text", "text")

	weekly := WeeklyTrackChart{}
	err = json.Unmarshal([]byte(valid), &weekly)
	if err != nil {
		return nil, err
	}

	return weekly.Tracks, nil
}
