package cli

import (
	"flag"

	"github.com/Gunyoung-Kim/info-join-test/info_join_test/tester"
)

// Start command line interface
func Start() {
	numOfRequest := flag.Int("numOfRequest", 1, "Set number of request made by this system")
	testURL := flag.String("url", "http://localhost:8080", "Set test URL")

	flag.Parse()
	tester.JoinTest(*testURL, *numOfRequest)
}
