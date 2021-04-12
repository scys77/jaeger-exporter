package app

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"sort"
	"strconv"
)

var (
	headerTraceID         = "Trace ID"
	headerSuffixStartTime = "Start Time"
	headerSuffixDuration  = "Duration"
)

// WriteToCSV writes the given Traces as a CSV to the filename
func WriteToCSV(traces []*Trace, filename string) error {
	// uniqueSpanNames := getUniqueSpanKeys(traces)

	// prepare file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// generate and write headers
	headers := []string{}

	headers = append(
		headers,
		"traceID",
		"spanID",
		"operationName",
		"process.serviceName",
		"references",
		"tags",
		"startTime",
		"duration",
	)

	// for _, operationName := range uniqueSpanNames {
	// 	headers = append(
	// 		headers,
	// 		operationName+" "+headerSuffixStartTime,
	// 		operationName+" "+headerSuffixDuration,
	// 	)
	// }

	err = writer.Write(headers)
	if err != nil {
		return err
	}

	// generate and write lines
	var line []string
	for _, resultItem := range traces {

		for _, span := range resultItem.Spans {
			line = []string{resultItem.TraceID}

			if span == nil {
				line = append(line, "", "")
				continue
			}

			reference, _ := json.Marshal(span.References)
			tags, _ := json.Marshal(span.Tags)

			line = append(line,
				// string(span.TraceID),
				string(span.SpanID),
				string(span.OperationName),
				string(span.SpanProcess.ServiceName),
				string(reference),
				string(tags),
				strconv.FormatUint(span.StartTime, 10),
				strconv.FormatUint(span.Duration, 10),
			)

			err = writer.Write(line)
			if err != nil {
				return err
			}
		}

		// for _, operationName := range uniqueSpanNames {
		// 	span := resultItem.Spans[operationName]
		// 	if span == nil {
		// 		line = append(line, "", "")
		// 		continue
		// 	}

		// 	line = append(line,
		// 		strconv.FormatUint(span.StartTime, 10),
		// 		strconv.FormatUint(span.Duration, 10),
		// 	)
		// }

	}

	return nil
}

func getUniqueSpanKeys(traces []*Trace) []string {
	var operationNames []string

	for _, trace := range traces {
		for name := range trace.Spans {

			if contains(operationNames, name) {
				continue
			}

			operationNames = append(operationNames, name)
		}
	}

	sort.Strings(operationNames)

	return operationNames
}

func contains(list []string, key string) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}

	return false
}
