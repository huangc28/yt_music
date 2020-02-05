package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SearchPlayListSuite struct {
	suite.Suite
}

func (suite *SearchPlayListSuite) TestSearchPlayListRoute() {
	// initialize a new request to search playlist API
	req, err := http.NewRequest("GET", "/search-playlist", nil)

	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("search", "hiphop")
	req.URL.RawQuery = q.Encode()

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchPlaylist)

	handler.ServeHTTP(rr, req)

	log.Printf("status code %v", rr.Code)
	log.Printf("Body %v", rr.Body.String())
}

// TestSearchPlayListWithoutCriteria test if 'search' is provided in the query param.
// if it is not provided it should return error with failed status code.
func (suite *SearchPlayListSuite) TestSearchPlayListWithoutCriteria() {
	req, err := http.NewRequest("GET", "/search-playlist", nil)

	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchPlaylist)
	handler.ServeHTTP(rr, req)

	// assert that the status code and the returned string is as expected

	assert.Equal(suite.T(), rr.Code, 400)
	log.Printf("Body %v", rr.Body.String())
}

func TestSearchPlayListSuite(t *testing.T) {
	suite.Run(t, new(SearchPlayListSuite))
}
