package lib

import (
	"log"
	"os"
	"path"
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

func CreateFile(filefolder, fileName, url string, size int64) (*os.File, error) {
	if f, err := os.Open(filefolder); err != nil {
		if f != nil {
			defer f.Close()
		}
		if os.IsNotExist(err) {
			log.Printf("download folder %s does not exist, create", filefolder)
			if err := os.MkdirAll(filefolder, 0755); err != nil {
				return nil, err
			}
		}
	}
	fullPath := path.Join(filefolder, fileName)
	size, err := GetStreamSize(url)
	if err != nil {
		return nil, err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	if err := f.Truncate(size); err != nil {
		return nil, err
	}
	return f, nil
}

func SliceWrite(size int64, f *os.File) error {
	//left := size
	//HttpDoGet()
	return nil
}

func SetDownloadFolder(folder string) {
	DownloadPath = folder
}
