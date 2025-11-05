package main

import (
	"fmt"

	"github.com/saltfishh/blDownLoader/lib"
)

func main() {
	//lib.SetCookie("SESSDATA=48439129%2C1777700085%2C9b13d%2Ab1CjDwK2eMNyncWTBhlD49V6ohXEb7_JNYFV2yUcSjrYjINlkbbnGMA1C9eEFSnDOG9kcSVnVtNEI1VHV5V0F0V3BYUTJfdEttaWlqdmhNZWtLLXl4ZThTVnpxLUhqc29PX3B1NnhPNzZETllMUlh5WWEtTmNOVG14SGVaRmFxbk9sVmE5dXFvQ25BIIEC")
	if err := lib.Downloader(""); err != nil {
		fmt.Println(err)
	}
}
