package lastfm

type GetTokenResponse struct {
	Token string `json:"token"`
}

type GetSessionResponse struct {
	Session struct{ Name, Key string } `json:"session"`
}
