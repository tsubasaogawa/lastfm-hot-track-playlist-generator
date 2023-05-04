package ytmusic

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const (
	TOKEN_FILE = "token.json"
)

func NewService(oauthfile string) (*youtube.Service, error) {
	oauth, err := os.ReadFile(oauthfile)
	if err != nil {
		return nil, err
	}

	cfg, err := google.ConfigFromJSON(oauth, youtube.YoutubeScope)
	if err != nil {
		return nil, err
	}

	cdir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	file := filepath.Join(cdir, "lfm2ytm", TOKEN_FILE)

	tok, err := readSavedToken(file)
	if err != nil {
		tok, err = generateToken(file, cfg)
		if err != nil {
			return nil, err
		}
		log.Println("Saving token file to " + file)
	}

	ctx := context.Background()
	c := cfg.Client(ctx, tok)
	s, err := youtube.New(c)

	return s, err
}

func readSavedToken(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := oauth2.Token{}
	err = json.NewDecoder(f).Decode(&tok)

	return &tok, err
}

func generateToken(file string, c *oauth2.Config) (*oauth2.Token, error) {
	hash := md5.Sum([]byte(time.Now().String()))
	url := c.AuthCodeURL(hex.EncodeToString(hash[:]), oauth2.AccessTypeOffline)
	fmt.Printf("Access to the following URL: \n%s\nAuth code is: ", url)

	var s string
	var sc = bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		s = sc.Text()
	}

	tok, err := c.Exchange(oauth2.NoContext, s)
	if err != nil {
		return nil, err
	}

	// save
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(tok)

	return tok, err
}
