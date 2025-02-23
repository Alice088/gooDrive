package gooDrive

import (
	"context"
	"golang.org/x/oauth2"
)

// GetClient Retrieve a token, saves the token, then returns the generated client.
func (drive *GooDrive) getClient(config *oauth2.Config) {
	tok, err := drive.tokenFromFile()
	if err != nil {
		tok = drive.getTokenFromWeb(config)
		drive.saveToken(tok)
	}
	drive.Client = config.Client(context.Background(), tok)
}
