package bilibili

type VideoStream struct {
	Quality           int      `json:"quality"`
	TimeLength        int      `json:"timelength"`
	AcceptDescription []string `json:"accept_description"`
	AcceptQuality     []int    `json:"accept_quality"`
	Dash              DashData `json:"dash"`
}
type DashData struct {
	Video []MediaData `json:"video"`
	Audio []MediaData `json:"audio"`
}
type MediaData struct {
	Id        int      `json:"id"`
	BaseUrl   string   `json:"baseUrl"`
	BackupUrl []string `json:"backupUrl"`
	FrameRate string   `json:"frameRate,omitempty"`
}

type VideoStreamData struct {
	Data *VideoStream `json:"data,omitempty"`
	ErrMsg
}
