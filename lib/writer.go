package lib

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func DownloadFile(filepath, fileName string, url string) error {

	/*
		f, err := os.Create(fullPath)
		if err != nil {
			if os.IsExist(err) {
				log.Printf("file %s already exists", fullPath)
			}
			return err
		}
		defer f.Close()
		resp, err := HttpDoGet(url, map[string]string{"referer": "https://www.bilibili.com"})
		if err != nil {
			return err
		}
		if i, err := f.Write(resp); err != nil {
			return err
		} else {
			log.Printf("download video %d bytes", i)
		}
	*/

	return nil
}

func CreateFile(fileName, url string) (*os.File, int64, error) {
	fullPath := path.Join(DownloadPath, fileName)
	size, err := GetStreamSize(url)
	if err != nil {
		return nil, 0, err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return nil, 0, err
	}
	if err := f.Truncate(size); err != nil {
		return nil, 0, err
	}
	return f, size, nil
}

func SetDownloadFolder(folder string) {
	DownloadPath = folder
}

func init() {
	if _, err := os.Stat(".env"); err != nil {
		if os.IsNotExist(err) {
			if f, err := os.Create(".env"); err != nil {
				log.Fatalf("create .env file failed: %v", err)
			} else {
				f.Close()
			}
		}
	}
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	biliSession := os.Getenv("BILI_SESSION")
	downloadFolder := os.Getenv("BILI_DOWNLOAD_FOLDER")
	if downloadFolder == "" {
		SetDownloadFolder("./download")
	} else {
		SetDownloadFolder(downloadFolder)
	}
	if biliSession != "" {
		SetCookie(biliSession)
	}
	log.Printf("download folder is %s", DownloadPath)
	if f, err := os.Open(downloadFolder); err != nil {
		if f != nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			log.Printf("download folder %s does not exist, create", DownloadPath)
			if err := os.MkdirAll(downloadFolder, 0755); err != nil {
			}
		}
	}
}
