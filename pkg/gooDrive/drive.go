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
	client    *http.Client
	service   *drive.Service
}

type IGooDrive interface {
	DownloadFile(fileId string, filePath FilePath) (FilePath, error)
	UploadFile(filePath string) (FileId, error)
	Client() *http.Client
	Service() *drive.Service
	getClient(config *oauth2.Config, tokensSavePath string)
	getTokenFromWeb(config *oauth2.Config) *oauth2.Token
	tokenFromFile(file string) (*oauth2.Token, error)
}

/*
NewGooDrive

tokenPath - This is the path to the file token and the path where the file token will be saved if it is not found
*/
func NewGooDrive(tokenPath string, tokensSavePath string) IGooDrive {
	ctx := context.Background()
	b, err := os.ReadFile(tokenPath)
	if err != nil {
		log.Fatalf("unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("unable to parse client secret file to config: %v", err)
	}

	gooDrive := &GooDrive{
		tokenPath: tokenPath,
	}

	gooDrive.getClient(config, tokensSavePath)

	gooDrive.service, err = drive.NewService(ctx, option.WithHTTPClient(gooDrive.client))
	if err != nil {
		log.Fatalf("unable to retrieve Drive client: %v", err)
	}

	return gooDrive
}
