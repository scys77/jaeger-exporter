package app

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	limit := flag.Int("limit", 30, "limit for the number of traces to fetch")
	flag.Parse()

	if *limit <= 0 {
		fmt.Println("The 'limit' parameter must be greater than 0.")
		os.Exit(1)
	}

	// Assuming the application has a function to handle the fetching of traces
	// This is a placeholder for where the application would use the 'limit' value
	// fetchTraces(*limit)
}
