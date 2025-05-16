package tests

import (
	"github.com/Alice088/gooDrive/pkg/env"
	"github.com/Alice088/gooDrive/pkg/gooDrive"
	"os"
	"testing"
)

func TestDownloadFileFromEnv(b *testing.T) {
	env.Init()

	filePath := "./session"
	drive := gooDrive.NewGooDriveENV("TOKEN", "../token.json")
	_, err := drive.DownloadFolder("1DR2h7NtrRaWcCGdKR9DkOhp1CYmlOpLn", filePath)

	if err != nil {
		b.Fatal(err)
	}

	_, err = os.Stat(filePath)

	if err != nil {
		b.Fatal(err)
	}

	err = os.RemoveAll(filePath)

	if err != nil {
		b.Fatal(err)
	}
}
