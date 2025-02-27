package gooDrive

import (
	googleDrive "google.golang.org/api/drive/v3"
	"log"
	"os"
	"path/filepath"
)

type FileId = string

func (drive *GooDrive) UploadFile(filePath string) (FileId, error) {
	f, err := os.Open(filePath)
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatalf("Error closing file during Google Drive upload: %v", err)
		}
	}(f)

	if err != nil {
		return "", err
	}

	driveFile := &googleDrive.File{
		Name: filepath.Base(filePath),
	}

	resp, err := drive.Service().Files.Create(driveFile).Media(f).Do()

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
