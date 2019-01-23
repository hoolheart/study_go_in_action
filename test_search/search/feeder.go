package search

// Feed contains information we need to process a feed.
type Feed struct {
	Name string
	URI  string
	Type string
}

// Feeder provides specific feeds
type Feeder interface {
	RetrieveFeeds() ([]*Feed, error)
}
