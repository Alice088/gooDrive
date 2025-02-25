package gooDrive

import "google.golang.org/api/drive/v3"

func (drive *GooDrive) Service() *drive.Service {
	return drive.service
}
