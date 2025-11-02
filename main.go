package main

import (
	"fmt"
	"log"

	"github.com/saltfishh/blDownLoader/lib"
)

func main() {
	vd, err := lib.GetVideoDetail("BV15AzhYSEbm")
	if err != nil {
		fmt.Println(err)
	}
	if vsd, err := lib.GetVideoStream(vd.Data.Aid, vd.Data.Cid); err != nil {
		fmt.Println(err)
	} else {
		if err := lib.DownloadFile("/home/wind/tmp/motrix/bldownload", vd.Data.Title+".video", vsd.Data.Dash.Video[0].BaseUrl); err != nil {
			fmt.Println(err)
		} else {
			log.Printf("video download: %s\n", vd.Data.Title)
		}

		if err := lib.DownloadFile("/home/wind/tmp/motrix/bldownload", vd.Data.Title+".audio", vsd.Data.Dash.Audio[0].BaseUrl); err != nil {
			fmt.Println(err)
		} else {
			log.Printf("audio download: %s\n", vd.Data.Title)
		}
	}
}
