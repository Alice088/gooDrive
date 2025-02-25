package gooDrive

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
)

// Request a token from the web, then returns the retrieved token.
func (drive *GooDrive) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	var authCode string

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("unable to retrieve token from web %v", err)
	}
	return tok
}
