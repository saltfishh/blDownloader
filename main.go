package main

import (
	"fmt"

	"github.com/saltfishh/blDownLoader/lib"
)

func main() {
	//url := "https://upos-sz-mirrorcos.bilivideo.com/upgcxcode/42/92/33696909242/33696909242-1-30112.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&deadline=1762528813&trid=33c31c175629480eb87eeab2b9b5884u&uipk=5&gen=playurlv3&oi=3730443642&nbs=1&platform=pc&os=cosbv&og=cos&mid=13979320&upsig=3a54f4d9db0267cd24e16a2bb6028f75&uparams=e,deadline,trid,uipk,gen,oi,nbs,platform,os,og,mid&bvc=vod&nettype=0&bw=2356868&agrr=0&buvid=&build=0&dl=0&f=u_0_0&orderid=1,3"
	if err := lib.Downloader("BV1Ng4y1K7CJ"); err != nil {
		fmt.Println(err)
	}
	//lib.HHH()
}
