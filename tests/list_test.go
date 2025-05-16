package tests

import (
	"github.com/Alice088/gooDrive/pkg/env"
	"github.com/Alice088/gooDrive/pkg/gooDrive"
	"testing"
)

func TestList(b *testing.T) {
	env.Init()
	drive := gooDrive.NewGooDriveENV("TOKEN", "../token.json")

	list, err := drive.Service().Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		b.Fatalf("Unable to retrieve files: %v", err)
	}
	b.Log("Files:")
	if len(list.Files) == 0 {
		b.Log("No files found.")
	} else {
		for _, i := range list.Files {
			b.Logf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
