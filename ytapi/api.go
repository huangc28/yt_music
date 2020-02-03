package ytapi

import (
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

type YoutubeApi struct {
	Service *youtube.Service
}

// @TODO client injected
func NewYoutubeAPI(devKey string) (*YoutubeApi, error) {
	// Initialize http client and youtube service
	client := &http.Client{
		Transport: &transport.APIKey{
			Key: devKey,
		},
	}

	service, err := youtube.New(client)

	if err != nil {
		return nil, err
	}

	return &YoutubeApi{
		Service: service,
	}, nil
}

// try requesting youtube search api for paylist result
func (ya *YoutubeApi) SearchPlaylist(criteria string) (*youtube.SearchListResponse, error) {
	call := ya.Service.Search.List("id,snippet").
		Q(queryCriteria).
		Type("playlist").
		MaxResults(*maxResults)

	return call.Do()
}
