package tests

import (
	"fmt"
	"github.com/Alice088/gooDrive/pkg/gooDrive"
	"log"
	"testing"
)

func TestNewGooDrive(b *testing.T) {
	drive := gooDrive.NewGooDrive("../credits.json", "../token.json")

	r, err := drive.Service().Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
