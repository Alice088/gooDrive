package gooDrive

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func (drive *GooDrive) saveToken(token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", drive.tokenPath)
	f, err := os.OpenFile(drive.tokenPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
