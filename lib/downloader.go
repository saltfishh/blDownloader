package lib

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var DownloadPath = ""

func Downloader(bvid string) error {
	vd, err := GetVideoDetail(bvid)
	if err != nil {
		return err
	}
	videoFile := strings.ReplaceAll(vd.Data.Title+".video", " ", "")
	audioFile := strings.ReplaceAll(vd.Data.Title+".audio", " ", "")
	vsd, err := GetVideoStream(vd.Data.Aid, vd.Data.Cid)
	if err != nil {
		return err
	}
	if err := SliceDownload(vsd.Data.Dash.Video[0].BaseUrl, videoFile); err != nil {
		return err
	}
	if err := SliceDownload(vsd.Data.Dash.Audio[0].BaseUrl, audioFile); err != nil {
		return err
	}
	if err := MergeAudioVideo(DownloadPath+videoFile, DownloadPath+audioFile, DownloadPath+vd.Data.Title+".mp4"); err != nil {
		return err
	}
	clearCache([]string{DownloadPath + videoFile, DownloadPath + audioFile})

	return nil
}

func clearCache(filePath []string) {
	for _, element := range filePath {
		if err := os.Remove(element); err != nil {
			log.Printf("Error deleting cache file: %v", err)
		}
	}
}
func SliceDownload(url, filename string) error {
	file, totalSize, err := CreateFile(filename, url)
	//log.Printf("total size: %d\n", totalSize)
	defer file.Close()
	if err != nil {
		log.Fatalf("%v", err)
	}
	m := make(map[string]string)
	m["referer"] = "https://www.bilibili.com"
	offset := int64(0)
	blockSize := int64(1024 * 1024 * 4)
	for offset < totalSize {
		m["Range"] = fmt.Sprintf("bytes=%d-%d", offset, offset+blockSize)
		rsp, err := HttpDoGet(url, m)
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Printf("download: %d to %d", offset, offset+blockSize)
		if _, err := file.WriteAt(rsp, offset); err != nil {
			return err
		}
		offset += blockSize
	}
	return nil
}
