package youtubeapi

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	queryCriteria = "hiphop"
	maxResults    = flag.Int64("max-results", 25, "Max YouTube results")
)

type YouTubeSDKSuite struct {
	suite.Suite
	service *youtube.Service
}

const developerKey = "AIzaSyD_E0RBoXGDvk68e7Dx-qwaW-RuOFptMzg"

func (suite *YouTubeSDKSuite) SetupTest() {
	// Initialize youtube client
	client := &http.Client{
		Transport: &transport.APIKey{
			Key: developerKey,
		},
	}

	service, err := youtube.New(client)

	if err != nil {
		log.Fatalf("Unable to initialize youtube client %v", err)
	}

	suite.service = service
}

func (suite *YouTubeSDKSuite) SearchPlayList() (*youtube.SearchListResponse, error) {
	call := suite.service.Search.List("id,snippet").
		Q(queryCriteria).
		Type("playlist").
		MaxResults(*maxResults)

	return call.Do()
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}

func (suite *YouTubeSDKSuite) TestGetPlayListFromYouTubeSDK() {
	// Make the API call to youtube
	resp, err := suite.SearchPlayList()

	assert.Nil(suite.T(), err)

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range resp.Items {
		// assert that the playlist ID for all items aren't empty string
		assert.NotEqual(suite.T(), "", item.Id.PlaylistId)
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	printIDs("Videos", videos)
	printIDs("Channels", channels)
	printIDs("Playlists", playlists)
}

func (suite *YouTubeSDKSuite) TestGetPlayListItemDetailFromPlayListID() {
	// Retrieve first item of the playlist
	resp, err := suite.SearchPlayList()

	assert.Nil(suite.T(), err)

	// Retrieve the first item from playlist
	playlistId := resp.Items[0].Id.PlaylistId

	// Use playlistId to retrieve playlist item
	call := suite.
		service.
		PlaylistItems.List("id,snippet").
		PlaylistId(playlistId).
		MaxResults(*maxResults)

	resp2, err := call.Do()

	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the first item from the playlist
	playlistItem := resp2.Items[0]

	assert.NotEqual(suite.T(), playlistItem.Snippet.ResourceId.VideoId, "")
}

func TestYouTubeSDKSuite(t *testing.T) {
	suite.Run(t, new(YouTubeSDKSuite))
}
