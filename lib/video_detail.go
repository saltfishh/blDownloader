package lib

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/saltfishh/blDownLoader/entity/bilibili"
	"github.com/saltfishh/blDownLoader/entity/consts"
)

func GetVideoDetail(bv string) (bilibili.VideoDetailsData, error) {
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
	if vd.Code != 0 {
		return vd, errors.New("fail to get video detail: " + vd.ErrMsg.Message)
	}
	//log.Printf("VideoDetails: %d", vd.Data.Owner.Mid)
	return vd, nil
}
