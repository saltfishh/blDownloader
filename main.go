package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/saltfishh/blDownLoader/lib"
)

func main() {
	vd, err := lib.GetVideoDetail("BV1D3411N7fs")
	if err != nil {
		fmt.Println(err)
	}
	if vsd, err := lib.GetVideoStream(vd.Data.Aid, vd.Data.Cid); err != nil {
		fmt.Println(err)
	} else {
		s, _ := json.Marshal(vsd)
		log.Printf("vsd: %s\n", s)
	}
}
