package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Push() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":       "54.184.220.131:9092",
		"broker.version.fallback": "0.10.0.0",
		"api.version.fallback.ms": 0,
		"sasl.mechanisms":         "PLAIN",
		"security.protocol":       "SASL_SSL",
		"sasl.username":           "AKIAQKC4VVTZLEREAGFG",
		"sasl.password":           "u5yYvrUdTVF8Z7oQW8/abptT6COWaKWfnyVaYMj"})
	if err != nil {
		panic(fmt.Sprintf("Failed to create producer: %s", err))
	}
	value := "golang test value"
	topic := "picture"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, nil)
	e := <-p.Events()

	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("failed to deliver message: %v\n", m.TopicPartition)
	} else {
		fmt.Printf("delivered to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	p.Close()
}
