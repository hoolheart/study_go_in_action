package search

import "log"

var feeders []Feeder
var matchers = make(map[string]Matcher)

// RegisterFeeder is called to append a feeder
func RegisterFeeder(feeder Feeder) {
	log.Println("Register a new feeder")
	feeders = append(feeders, feeder)
}

// RegisterMatcher is called to register a matcher for use by the program.
func RegisterMatcher(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
