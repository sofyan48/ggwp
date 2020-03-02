package kafka

import (
	"os"
	"time"

	"github.com/Shopify/sarama"
)

// Kafka ...
type Kafka struct{}

// KafkaHandler ...
func KafkaHandler() *Kafka {
	return &Kafka{}
}

// KafkaInterface ...
type KafkaInterface interface {
	Producer(topic string, msg string) (int64, error)
}

// Producer ...
func (kf *Kafka) Producer(topic string, msg string) (int64, error) {
	kafkaConfig := getKafkaConfig(os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	producer, err := sarama.NewSyncProducer([]string{kafkaHost + ":" + kafkaPort}, kafkaConfig)
	if err != nil {
		return 0, err
	}
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	_, offset, err := producer.SendMessage(kafkaMsg)
	if err != nil {
		return offset, err
	}
	return offset, nil
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
