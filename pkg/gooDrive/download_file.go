package gooDrive

import (
	"errors"
	"io"
	"log"
	"os"
)

type FilePath = string

func (drive *GooDrive) DownloadFile(fileId string, filePath FilePath) (FilePath, error) {
	resp, err := drive.Service().Files.Get(fileId).Download()

	if err != nil {
		return filePath, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(resp.Body)

	outFile, err := os.Create(filePath)
	if err != nil {
		return filePath, err

	}
	defer func(outFile *os.File) {
		err = outFile.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}(outFile)

	_, err = io.Copy(outFile, resp.Body)

	if errors.Is(err, io.EOF) {
		return filePath, nil
	}

	return filePath, err
}
