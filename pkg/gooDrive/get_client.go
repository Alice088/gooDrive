package gooDrive

import (
	"context"
	"golang.org/x/oauth2"
)

// GetClient Retrieve a token, saves the token, then returns the generated client.
func (drive *GooDrive) getClient(config *oauth2.Config, tokensSavePath string) {
	tokFile := tokensSavePath
	tok, err := drive.tokenFromFile(tokFile)
	if err != nil {
		tok = drive.getTokenFromWeb(config)
		drive.saveToken(tokFile, tok)
	}
	drive.client = config.Client(context.Background(), tok)
}
