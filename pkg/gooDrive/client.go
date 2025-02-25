package gooDrive

import "net/http"

func (drive *GooDrive) Client() *http.Client {
	return drive.client
}
