package gooDrive

import (
	"context"
	"github.com/Alice088/gooDrive/pkg/env"
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
NewGooDriveJSON

tokenPath - This is the path to the file token and the path where the file token will be saved if it is not found
*/
func NewGooDriveJSON(tokenPath string, tokensSavePath string) IGooDrive {
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

/*
NewGooDriveENV

name - Name of ENV value where is credits

savePath - Path to save token.json
*/
func NewGooDriveENV(name, savePath string) IGooDrive {
	credits, err := env.MuGet[[]byte](name)
	if err != nil {
		log.Fatal(err.Error())
	}

	// If modifying these scopes, delete your previously saved credits.json.
	config, err := google.ConfigFromJSON(credits, drive.DriveScope)
	if err != nil {
		log.Fatalf("unable to parse client secret file to config: %v", err)
	}

	gooDrive := &GooDrive{
		tokenPath: savePath,
	}

	gooDrive.getClient(config, savePath)

	ctx := context.Background()
	gooDrive.service, err = drive.NewService(ctx, option.WithHTTPClient(gooDrive.client))
	if err != nil {
		log.Fatalf("unable to retrieve Drive client: %v", err)
	}

	return gooDrive
}
