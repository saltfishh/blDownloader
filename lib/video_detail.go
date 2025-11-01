package lib

import (
	"encoding/json"
	"log"

	"github.com/saltfishh/blDownLoader/entity/bilibili"
	"github.com/saltfishh/blDownLoader/entity/consts"
)

func GetVideoDetail(bv string) bilibili.VideoDetailsData {
	vd := bilibili.VideoDetailsData{}
	burl, err := UrlConvert(consts.VideoDetailsUrl, map[string]string{"bvid": bv})
	if err != nil {
		log.Panicf("VideoDetailUrlConvert: %s", err)
	}
	rsp, err := HttpDoGet(burl, nil)
	if err != nil {
		log.Panicf("httpDoGet: %s", err)
	}
	if err := json.Unmarshal(rsp, &vd); err != nil {
		log.Printf("Unmarshal: %s", err)
	}
	//log.Printf("VideoDetails: %d", vd.Data.Owner.Mid)
	return vd
}
