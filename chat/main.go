package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sofyan48/ggwp/chat/consumer"
	"github.com/sofyan48/ggwp/chat/send"
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
	window := flag.String("w", "window", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	ConfigEnvironment(*environment)
	if *window == "send" {
		startChatting()
	} else {
		windowChat()
	}
}

func startChatting() {
	sendHandler := send.SendHandler()
	var ID string
	var room string
	fmt.Printf("Room Name: ")
	fmt.Scanln(&room)
	fmt.Printf("ID Name: ")
	fmt.Scanln(&ID)
	var chatValue, toValue string
	for {
		fmt.Println("------------------------------------------------")
		fmt.Println("Window Chat")
		fmt.Println("------------------------------------------------")
		fmt.Print("to : ")
		fmt.Scanln(&toValue)
		fmt.Print("Chat : ")
		fmt.Scanln(&chatValue)
		fmt.Println("------------------------------------------------")
		fmt.Println("Messages Sending")
		fmt.Println("Press Enter To Next")
		fmt.Println("------------------------------------------------")
		sendHandler.Send(room, ID, toValue, chatValue)
	}
}

func windowChat() {
	fmt.Println("------------------------------------------------")
	fmt.Println("Chat")
	fmt.Println("------------------------------------------------")
	var ID string
	var room string
	fmt.Printf("Room Name: ")
	fmt.Scanln(&room)
	os.Setenv("ROOM_KEY", room)
	fmt.Printf("ID Name: ")
	fmt.Scanln(&ID)
	os.Setenv("CHAT_ID", ID)
	initWorker(room, ID)
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
