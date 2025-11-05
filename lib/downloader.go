package lib

import (
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
	if err := DownloadFile(DownloadPath, videoFile, vsd.Data.Dash.Video[0].BaseUrl); err != nil {
		return err
	}
	if err := DownloadFile(DownloadPath, audioFile, vsd.Data.Dash.Audio[0].BaseUrl); err != nil {
		return err
	}
	/*	if err := MergeAudioVideo(DownloadPath+videoFile, DownloadPath+audioFile, DownloadPath+vd.Data.Title+".mp4"); err != nil {
			return err
		}
		clearCache([]string{DownloadPath + videoFile, DownloadPath + audioFile})
	*/
	return nil
}

func clearCache(filePath []string) {
	for _, element := range filePath {
		if err := os.Remove(element); err != nil {
			log.Printf("Error deleting cache file: %v", err)
		}
	}
}
