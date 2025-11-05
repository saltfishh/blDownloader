package lib

import (
	"io"
	"net/url"
	"strconv"

	"github.com/astaxie/beego/httplib"
)

var SessionDataHeader = make(map[string]string)

func HttpDoGet(targetUrl string, header map[string]string) ([]byte, error) {
	req := httplib.Get(targetUrl)
	req.Header("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")
	if header != nil {
		for k, v := range header {
			req.Header(k, v)
		}
	}
	rsp, err := req.DoRequest()
	if err != nil {
		//log.Printf("req.DoRequest err: %s", err)
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := io.ReadAll(rsp.Body)
	if err != nil {
		//log.Printf("io read rsp body err: %s", err)
		return nil, err
	}
	return b, err
}

func UrlConvert(baseUrl string, queryParm map[string]string) (string, error) {
	vdUrl, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	q := vdUrl.Query()
	if queryParm != nil {
		for k, v := range queryParm {
			q.Set(k, v)
		}
	}
	vdUrl.RawQuery = q.Encode()
	return vdUrl.String(), nil
}

func SetCookie(sessionData string) {
	SessionDataHeader["Cookie"] = sessionData
}

func GetStreamSize(url string) (int64, error) {
	req := httplib.Head(url)
	req.Header("referer", "http://www.bilibili.com")
	req.Header("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36")
	rsp, err := req.DoRequest()
	if err != nil {
		return 0, err
	}
	size := rsp.Header.Get("Content-Length")
	if size == "0" {
		return 0, nil
	}
	i, err := strconv.Atoi(size)
	if err != nil {
		return 0, err
	}
	return int64(i), nil
}
