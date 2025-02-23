package gooDrive

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

// GetClient Retrieve a token, saves the token, then returns the generated client.
func (drive *GooDrive) GetClient(config *oauth2.Config) *http.Client {
	tok, err := drive.tokenFromFile()
	if err != nil {
		tok = drive.getTokenFromWeb(config)
		drive.saveToken(tok)
	}
	return config.Client(context.Background(), tok)
}
