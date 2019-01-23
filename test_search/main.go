package main

import (
	"log"
	"os"
	"github.com/hoolheart/study_go_in_action/test_search/search"
	_ "github.com/hoolheart/study_go_in_action/test_search/feeders"
	_ "github.com/hoolheart/study_go_in_action/test_search/matchers"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
