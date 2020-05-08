package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sofyan48/ggwp/client/consumer"
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

	var ID string
	var room string
	fmt.Printf("Room Name: ")
	fmt.Scanln(&room)
	fmt.Printf("ID Name: ")
	fmt.Scanln(&ID)
	go initWorker(room, ID)

	var chatValue, toValue string
	for {
		fmt.Printf("to : ")
		fmt.Scanln(&toValue)
		fmt.Println(toValue)
		fmt.Printf("Chat : ")
		fmt.Scanln(&chatValue)
		fmt.Println(chatValue)
	}
}

func initWorker(room, ID string) {
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
	kafkaConsumer.Consume([]string{"chat"}, signals, room, ID)
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
