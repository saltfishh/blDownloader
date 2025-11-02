package lib

import (
	"encoding/json"
	"strconv"

	"github.com/saltfishh/blDownLoader/entity/bilibili"
	"github.com/saltfishh/blDownLoader/entity/consts"
)

func GetVideoStream(avid, cid int) (bilibili.VideoStreamData, error) {
	vsd := bilibili.VideoStreamData{}
	m := make(map[string]string)
	m["avid"] = strconv.Itoa(avid)
	m["cid"] = strconv.Itoa(cid)
	m["qn"] = strconv.Itoa(120)
	m["fnver"] = strconv.Itoa(0)
	m["fourk"] = strconv.Itoa(1)
	m["fnval"] = strconv.Itoa(80)
	vs, err := UrlConvert(consts.VideoStreamUrl, m)
	if err != nil {
		return vsd, err
	}
	//log.Printf("vs: %s\n", vs)
	rsp, err := HttpDoGet(vs, m)
	if err != nil {
		return vsd, err
	}
	if err := json.Unmarshal(rsp, &vsd); err != nil {
		return bilibili.VideoStreamData{}, err
	}
	return vsd, nil
}
