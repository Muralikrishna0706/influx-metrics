package producer

import (
	"context"

	"fmt"

	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InsertData() {

	// insert sample data into influx db

	// vendor code,RC,ARC,status, module

	token := "bOBfnE-sluaC2fXYPlbVPvKve1OtqFo5LGe-OhPymbYUfwo68YurwIGS6YIcwpJkePxucqjH-RS_oKvGF68vlQ=="

	client := influxdb2.NewClient("http://localhost:8086", token)

	writeAPI := client.WriteAPIBlocking("ayopop", "controlroom")

	p := influxdb2.NewPoint("stat",

		map[string]string{

			"module": "travel-loka",
		},

		map[string]interface{}{

			"vendor_code": "200",
			"RC":          "000",
			"ARC":         "200",
			"status":      "2",
			"module":      "inquiry",
		},

		time.Now())

	// write point immediately

	writeAPI.WritePoint(context.Background(), p)

	// create point using fluent style

	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("module", "travel-loka").
		SetTime(time.Now())

	writeAPI.WritePoint(context.Background(), p)

	// Or write directly line protocol

	line := fmt.Sprintf("stat,unit=code avg=%f,max=%f", 23.5, 45.0)

	writeAPI.WriteRecord(context.Background(), line)

	// Ensures background processes finish

	client.Close()

}
