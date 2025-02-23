package gooDrive

import (
	"golang.org/x/oauth2"
	"net/http"
)

type GooDrive struct {
	tokenPath string
}

type IGooDrive interface {
	GetClient(config *oauth2.Config) *http.Client
	getTokenFromWeb(config *oauth2.Config) *oauth2.Token
	tokenFromFile(file string) (*oauth2.Token, error)
}

/*
NewGooDrive

	tokenPath - there
*/
func NewGooDrive(tokenPath string) *GooDrive {

	return &GooDrive{
		tokenPath: tokenPath,
	}
}
