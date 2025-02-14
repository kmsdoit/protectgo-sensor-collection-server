package influxdb

import influxdb2 "github.com/influxdata/influxdb-client-go/v2"

const INFLUXDB_TOKEN = "2BxmRAS1OwpeICA8XdzvPU1TZBwKn9sNJJOpzwRjyR0qPoasIToMRgh6BEWM_BOa7JIIBI5f_XcBWMXKRCdqcw=="

func Init() influxdb2.Client {
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, INFLUXDB_TOKEN)

	return client
}
