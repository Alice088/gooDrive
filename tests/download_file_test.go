package tests

import (
	"github.com/Alice088/gooDrive/pkg/gooDrive"
	"os"
	"testing"
)

func TestDownloadFile(b *testing.T) {
	filePath := "./test_photo.png"
	drive := gooDrive.NewGooDrive("../credits.json", "../token.json")
	_, err := drive.DownloadFile("1CijSxpbBt-QeMdFh8JILAx12GOYtCiCT", filePath)

	if err != nil {
		b.Fatal(err)
	}

	_, err = os.Stat(filePath)

	if err != nil {
		b.Fatal(err)
	}

	err = os.Remove(filePath)

	if err != nil {
		b.Fatal(err)
	}
}
