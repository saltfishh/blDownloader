package lib

import (
	"log"
	"os"
	"path"
)

func DownloadFile(filepath, fileName string, url string) error {
	if f, err := os.Open(filepath); err != nil {
		if f != nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			log.Printf("download folder %s does not exist, create", filepath)
			if err := os.MkdirAll(filepath, 0755); err != nil {
				return err
			}
		}
	}
	fullPath := path.Join(filepath, fileName)
	f, err := os.Create(fullPath)
	if err != nil {
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
	return nil
}
