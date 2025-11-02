package lib

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/saltfishh/blDownLoader/entity/bilibili"
	"github.com/saltfishh/blDownLoader/entity/consts"
)

func GetVideoStream(avid, cid int) (bilibili.VideoStreamData, error) {
	vsd := bilibili.VideoStreamData{}
	m := make(map[string]string)
	m["avid"] = strconv.Itoa(avid)
	m["cid"] = strconv.Itoa(cid)
	m["qn"] = strconv.Itoa(160)
	m["fnver"] = strconv.Itoa(0)
	m["fourk"] = strconv.Itoa(1)
	m["fnval"] = strconv.Itoa(80)
	vs, err := UrlConvert(consts.VideoStreamUrl, m)
	if err != nil {
		return vsd, err
	}
	rsp, err := HttpDoGet(vs, SessionDataHeader)
	if err != nil {
		return vsd, err
	}
	if err := json.Unmarshal(rsp, &vsd); err != nil {
		return bilibili.VideoStreamData{}, err
	}
	if vsd.Code != 0 {
		return bilibili.VideoStreamData{}, errors.New("fetch stream error: " + vsd.ErrMsg.Message)
	}
	return vsd, nil
}
