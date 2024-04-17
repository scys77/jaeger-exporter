package main

import (
	"flag"
	"os"
	"fmt"
	"net/http"
	"time"

	"github.com/scys77/jaeger-exporter/pkg/app"
)

var (
	host     = ""
	limit    = 0
	lookback = ""
	service  = ""
	tags     = ""
	filename = ""
	username = ""
	password = ""
)

func main() {
	// parse flags
	flag.StringVar(&host, "host", "", "host of Jaeger, eg https://jaeger-query.company.com")
	flag.IntVar(&limit, "limit", 30, "maximum number of items, eg 30")
	flag.StringVar(&lookback, "lookback", "2d", "timerange, eg 2d")
	flag.StringVar(&service, "service", "", "name of the service, eg example-service")
	flag.StringVar(&tags, "tags", "", "tags to filter for, eg {\"foo\":\"bar\"}")
	flag.StringVar(&filename, "filename", "output.csv", "filename to write to, eg output.csv")
	flag.StringVar(&username, "username", "", "username for http basic auth")
	flag.StringVar(&password, "password", "", "password for http basic auth")
	flag.Parse()

	if limit <= 0 {
		fmt.Println("The 'limit' parameter must be greater than 0.")
		os.Exit(1)
	}

	// prepare http.Client
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// get data from Jaeger
	jaegerResp, err := app.JaegerSearchTraces(
		client,
		host,
		limit,
		lookback,
		service,
		tags,
		username,
		password,
	)
	if err != nil {
		panic(err)
	}

	// generate traces list with spans map
	traces := app.GenerateTraceWithSpansMap(jaegerResp)

	// write to CSV
	err = app.WriteToCSV(traces, filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d trace(s) to %s\n", len(jaegerResp.Data), filename)
}
