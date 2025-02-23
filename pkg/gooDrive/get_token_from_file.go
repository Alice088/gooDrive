package gooDrive

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"os"
)

// Request a token from the web, then returns the retrieved token.
func (drive *GooDrive) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
