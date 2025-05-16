package gooDrive

import (
	"fmt"
	"os"
	"path/filepath"
)

type FolderPath = string

func (drive *GooDrive) DownloadFolder(FolderId string, folderPath FolderPath) (FolderPath, error) {
	files, err := drive.service.Files.List().Q(fmt.Sprintf("'%s' in parents", FolderId)).Do()
	if err != nil {
		return folderPath, err
	}

	for _, file := range files.Files {
		filePath := filepath.Join(folderPath, file.Name)
		if file.MimeType == "application/vnd.google-apps.folder" {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return folderPath, err
			}
			_, err = drive.DownloadFolder(file.Id, filePath)
			if err != nil {
				return folderPath, err
			}
		} else {
			_, err := drive.DownloadFile(file.Id, filePath)
			if err != nil {
				return folderPath, err
			}
		}
	}

	return folderPath, nil
}
