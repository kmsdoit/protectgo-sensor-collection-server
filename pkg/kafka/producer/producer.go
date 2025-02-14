package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/kmsdoit/protectgo-sensor-collection-server/pkg/storage/influxdb"
	"math/rand"
	"time"
)

func ProduceSensorData(producer *kafka.Producer, inf_client influxdb2.Client, sensorType string) {
	topic := fmt.Sprintf("sensor-%s", sensorType)

	for {
		value := 20 + rand.Float64()*10 // 예제 데이터
		timestamp := time.Now().Unix()
		message := fmt.Sprintf(`{"sensor_type": "%s", "value": %.2f, "timestamp": %d}`,
			sensorType, value, timestamp)

		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)

		if err != nil {
			fmt.Println("Error producing message:", err)
		} else {
			fmt.Printf("Produced: %s -> %s\n", sensorType, message)
			influxdb.WriteToInfluxdb(inf_client, sensorType, value, timestamp)
		}

		time.Sleep(1 * time.Second) // 2초마다 센서 데이터 전송
	}
}
