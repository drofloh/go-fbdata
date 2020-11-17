package fbdata

import "net/http"

const (
	// BaseAPIURL ...
	BaseAPIURL = "http://api.football-data.org/v2"
)

type CLI struct {
	client *http.Client
	token  string
}
