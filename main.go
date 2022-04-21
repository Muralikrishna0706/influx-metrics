package main

import (
	"context"
	"influx-metrics/producer"

	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {

	// produce some sample data

	producer.InsertData()

	// Create client

	token := "bOBfnE-sluaC2fXYPlbVPvKve1OtqFo5LGe-OhPymbYUfwo68YurwIGS6YIcwpJkePxucqjH-RS_oKvGF68vlQ=="

	client := influxdb2.NewClient("http://localhost:8086", token)

	// Get query client

	queryAPI := client.QueryAPI("ayopop")

	// Get QueryTableResult

	query := `from(bucket:"controlroom")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`

	// get QueryTableResult

	result, err := queryAPI.Query(context.Background(), query)

	if err != nil {

		panic(err)

	}

	// Iterate over query response

	for result.Next() {

		// Notice when group key has changed

		if result.TableChanged() {

			fmt.Printf("table: %s\n", result.TableMetadata().String())

		}

		// Access data

		fmt.Printf("value: %v\n", result.Record().Value())

	}

	// check for an error

	if result.Err() != nil {

		fmt.Printf("query parsing error: %\n", result.Err().Error())

	}

}
