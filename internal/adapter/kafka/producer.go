package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

var Writer *kafka.Writer

func InitKafkaWriter() {
	Writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{os.Getenv("KAFKA_ADDR")},
		Topic:    "vacancy_created",
		Balancer: &kafka.LeastBytes{},
	})
}

func ProduceMessage(value []byte) {
	if Writer == nil {
		log.Println("Kafka Writer is nil") // критичный лог
		return
	}

	err := Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(time.Now().Format(time.RFC3339)),
			Value: value,
		},
	)
	if err != nil {
		log.Println("kafka produce error:", err)
	}
}
