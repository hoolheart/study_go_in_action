package search

import (
	"log"
	"sync"
)

var waitGroup sync.WaitGroup

func init() {
	waitGroup.Add(1)
}

// RetrieveFeeds retrieve all feeds from feeders, and trigger matching
func RetrieveFeeds(searchTerm string, results chan<- *Result) {
	for _, feeder := range feeders {
		feeds, err := feeder.RetrieveFeeds()
		if err == nil {
			log.Printf("Retrieving %d data\n",len(feeds))
			for _, feed := range feeds {
				// Retrieve a matcher for the search.
				matcher, exists := matchers[feed.Type]
				if !exists {
					matcher = matchers["default"]
				}
		
				// Launch the goroutine to perform the search.
				waitGroup.Add(1)
				go func(matcher Matcher, feed *Feed) {
					Match(matcher, feed, searchTerm, results)
					waitGroup.Done()
				}(matcher, feed)
			}
		} else {
			log.Printf("Failed to retrieve data: %s\n", err)
		}
	}

	log.Println("Finish retrieving")
	waitGroup.Done()
}

// MonitorSearching monitors all tasks in searching
func MonitorSearching(results chan *Result) {
	// Wait for everything to be processed.
	waitGroup.Wait()

	// Close the channel to signal to the Display
	// function that we can exit the program.
	close(results)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	log.Println(feed.URI)
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the individual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
