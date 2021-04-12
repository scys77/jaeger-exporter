# jaeger-exporter
Export Jaeger Traces to CSV.

## Installing:
```
go get github.com/scys77/jaeger-exporter/cmd/jaeger-exporter
go install github.com/scys77/jaeger-exporter/cmd/jaeger-exporter
```

## Usage:
```
jaeger-exporter

  -filename string
    	filename to write to, eg output.csv (default "output.csv")
  -host string
    	host of Jaeger, eg https://jaeger-query.company.com
  -limit int
    	maximum number of items, eg 20 (default 20)
  -lookback string
    	timerange, eg 2d (default "2d")
  -service string
    	name of the service, eg example-service
  -tags string
    	tags to filter for, eg {"foo":"bar"}
```


## local

```
cd cmd/jaeger-exporter/
go install

cd ~
./go/bin/jaeger-exporter

    -filename string
        filename to write to, eg output.csv (default "output.csv")
    -host string
        host of Jaeger, eg https://jaeger-query.company.com
    -limit int
        maximum number of items, eg 20 (default 20)
    -lookback string
        timerange, eg 2d (default "2d")
    -password string
        password for http basic auth
    -service string
        name of the service, eg example-service
    -tags string
        tags to filter for, eg {"foo":"bar"}
    -username string
        username for http basic auth
```
