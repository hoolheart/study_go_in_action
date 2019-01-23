package feeders

import (
	"encoding/json"
	"os"
	"github.com/hoolheart/study_go_in_action/test_search/search"
)

const dataFile = "feeders/data.json"

// jsonFeed contains information we need to process a feed by using json marking
type jsonFeed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

type jsonFeeder struct {}

func init() {
	var feeder jsonFeeder
	search.RegisterFeeder(feeder)
}

func (f jsonFeeder)RetrieveFeeds() ([]*search.Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()

	// Decode the file into a slice of pointers
	// to Feed values.
	var jsonFeeds []*jsonFeed
	err = json.NewDecoder(file).Decode(&jsonFeeds)

	// Convert to search Feed type
	var feeds []*search.Feed
	for _, feed := range jsonFeeds {
		var sFeed search.Feed
		sFeed.Name = feed.Name//copy
		sFeed.URI = feed.URI
		sFeed.Type = feed.Type
		feeds = append(feeds, &sFeed)//append into array
	}

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
