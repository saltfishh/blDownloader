package bilibili

type VideoStream struct {
}

type videoDetails struct {
	Aid    int    `json:"aid"`
	Videos int    `json:"videos"`
	Title  string `json:"title"`
	Stat   struct {
		View int `json:"view"`
	} `json:"stat"`
	Cid   int `json:"cid"`
	Owner struct {
		Mid  int    `json:"mid"`
		Name string `json:"name"`
	} `json:"owner"`
}
type VideoDetailsData struct {
	Data *videoDetails `json:"data,omitempty"`
	ErrMsg
}

type ErrMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
