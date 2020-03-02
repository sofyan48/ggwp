package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sofyan48/ggwp/worker/consumer"
)

// ConfigEnvironment |
func ConfigEnvironment(env string) {
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			logrus.Fatal("Error loading .env file")
		}
	}
}

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	ConfigEnvironment(*environment)
	initWorker()
}

func initWorker() {
	kafkaConfig := getKafkaConfig(os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	consumers, err := sarama.NewConsumer([]string{kafkaHost + ":" + kafkaPort}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Error create kakfa consumer got error %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			logrus.Fatal(err)
			return
		}
	}()

	kafkaConsumer := &consumer.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	kafkaConsumer.Consume([]string{"test_topic"}, signals)
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
