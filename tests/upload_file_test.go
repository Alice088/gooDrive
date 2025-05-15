package tests

import (
	"github.com/Alice088/gooDrive/pkg/gooDrive"
	"testing"
)

func TestUploadFile(b *testing.T) {
	drive := gooDrive.NewGooDriveJSON("../credits.json", "../token.json")
	fId, err := drive.UploadFile("../internal/photo_1.jpg")

	if err != nil {
		b.Fatal(err)
	}

	b.Log(fId)

	err = drive.Service().Files.Delete(fId).Do()

	if err != nil {
		b.Logf("Error deleting file: %s \n: %v", fId, err)
	} // /
}
