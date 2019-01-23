package search

// Run performs the search logic.
func Run(searchTerm string) {
	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Start Retrieving feeds
	go RetrieveFeeds(searchTerm,results)

	//Start monitoring
	go MonitorSearching(results)

	//Display result
	Display(results)
}
