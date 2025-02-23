package gooDrive

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

type GooDrive struct {
	tokenPath string
	Client    *http.Client
	Service   *drive.Service
}

type IGooDrive interface {
	getClient(config *oauth2.Config)
	getTokenFromWeb(config *oauth2.Config) *oauth2.Token
	tokenFromFile() (*oauth2.Token, error)
}

/*
NewGooDrive

tokenPath - This is the path to the file token and the path where the file token will be saved if it is not found
*/
func NewGooDrive(tokenPath string) *GooDrive {
	ctx := context.Background()
	b, err := os.ReadFile(tokenPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	gooDrive := &GooDrive{
		tokenPath: tokenPath,
	}

	gooDrive.getClient(config)

	gooDrive.Service, err = drive.NewService(ctx, option.WithHTTPClient(gooDrive.Client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	return gooDrive
}
