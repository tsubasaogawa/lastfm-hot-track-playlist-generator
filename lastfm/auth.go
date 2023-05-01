package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	apikey  string
	secret  string
	token   string
	session string
}

func NewAuth(key, sec string) (*Auth, error) {
	tok, err := getToken(key)
	if err != nil {
		return nil, err
	}

	return &Auth{
		apikey: key,
		secret: sec,
		token:  tok,
	}, nil
}

func getToken(key string) (string, error) {
	url := fmt.Sprintf("%s?method=auth.getToken&api_key=%s&format=json",
		ENDPOINT_BASE, key)
	_resp, err := request(url)
	if err != nil {
		return "", err
	}

	resp := GetTokenResponse{}
	err = json.Unmarshal(_resp, &resp)
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func request(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (auth *Auth) signature(method string) string {
	plain := fmt.Sprintf("api_key%smethod%stoken%s%s", auth.apikey, method, auth.token, auth.secret)
	hash := md5.Sum([]byte(plain))

	return hex.EncodeToString(hash[:])
}

func (auth *Auth) CreateGrantUrl() string {
	return fmt.Sprintf("https://www.last.fm/api/auth/?api_key=%s&token=%s", auth.apikey, auth.token)
}

func (auth *Auth) SetSessionKey() error {
	sig := auth.signature("auth.getSession")
	url := fmt.Sprintf("%s?method=auth.getSession&api_key=%s&api_sig=%s&token=%s&format=json",
		ENDPOINT_BASE, auth.apikey, sig, auth.token)

	_resp, err := request(url)
	if err != nil {
		return err
	}
	resp := GetSessionResponse{}
	err = json.Unmarshal(_resp, &resp)
	if err != nil {
		return err
	}

	auth.session = resp.Session.Key

	return nil
}
