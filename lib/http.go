package lib

import (
	"io"
	"log"
	"net/url"

	"github.com/astaxie/beego/httplib"
)

func HttpDoGet(targetUrl string, header map[string]string) ([]byte, error) {
	req := httplib.Get(targetUrl)
	if header != nil {
		for k, v := range header {
			req.Header(k, v)
		}
	}
	rsp, err := req.DoRequest()
	if err != nil {
		log.Printf("req.DoRequest err: %s", err)
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("io read rsp body err: %s", err)
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
