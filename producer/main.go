package main

import (
	"fmt"
	"github.com/protectgo-sensor-collection-server/internal/kafka"
)

func main() {
	fmt.Println("producer started")
	sensor := []string{
		"radionode",
	}
	kafka.ProducerInit(sensor)
}
