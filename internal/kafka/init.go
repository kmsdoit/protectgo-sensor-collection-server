package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	producer2 "github.com/kmsdoit/protectgo-sensor-collection-server/internal/kafka/producer"
	"github.com/kmsdoit/protectgo-sensor-collection-server/internal/storage/influxdb"
	"sync"
)

func ProducerInit(sensors []string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	inf_client := influxdb.Init()

	if err != nil {
		panic("Kafka Producer Error: " + err.Error())
	}
	defer producer.Close()

	var wg sync.WaitGroup

	for _, sensor := range sensors {
		wg.Add(1)
		go func(sensor string) {
			defer wg.Done()
			producer2.ProduceSensorData(producer, inf_client, sensor)
		}(sensor)
	}

	wg.Wait()
}
