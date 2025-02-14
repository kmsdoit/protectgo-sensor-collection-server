package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

func WriteToInfluxdb(client influxdb2.Client, sensorType string, value float64, timestamp int64) {
	writeAPI := client.WriteAPIBlocking("idb", "sensor")
	tags := map[string]string{
		"tag": sensorType,
	}
	fields := map[string]interface{}{
		"value":     value,
		"timestamp": timestamp,
	}
	point := write.NewPoint("radionode", tags, fields, time.Now())
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		log.Fatal(err)
	}
}
