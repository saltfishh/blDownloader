package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/saltfishh/blDownLoader/lib"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	biliSession := os.Getenv("BILI_SESSION")
	downloadFolder := os.Getenv("BILI_DOWNLOAD_FOLDER")
	if downloadFolder == "" {
		lib.SetDownloadFolder("./download")
	} else {
		lib.SetDownloadFolder(downloadFolder)
	}
	if biliSession != "" {
		lib.SetCookie(biliSession)
	}
	if err := lib.Downloader("BV1rk4y1j7o5"); err != nil {
		fmt.Println(err)
	}

}
