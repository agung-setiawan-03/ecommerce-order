package cmd

import (
	"ecommerce-order/helpers"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
)

func ServeKafkaConsumer() {
	brokers := strings.Split(helpers.GetEnv("KAFKA_HOST", "localhost:9092"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC", "example-topic")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Failed to connect with kafka as consumer", err)
		return
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		helpers.Logger.Error("Failed to create consume partition 01", err)
		return
	}

	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s from partition %d\n", string(msg.Value), msg.Partition)
	}
}