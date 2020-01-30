package youtubeapi

import "net/http"

type Youtuber struct{}

func GetYoutuber(cli *http.Client) *Youtuber {
	return &Youtuber{}
}

// try requesting youtube search api for paylist result
func GetPlayListByQuery() string {
	return "hello query"
}
