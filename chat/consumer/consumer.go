package consumer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/sofyan48/ggwp/chat/entity"
)

// KafkaConsumer hold sarama consumer
type KafkaConsumer struct {
	Consumer sarama.Consumer
}

// Consume function to consume message from apache kafka
func (c *KafkaConsumer) Consume(topics []string, signals chan os.Signal, room, ID string) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	for _, topic := range topics {
		partitionList, err := c.Consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, topic, partition, chanMessage)
		}
	}

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			payloadData := &entity.PayloadStateFull{}
			json.Unmarshal(msg.Value, payloadData)
			hearChat(room, ID, payloadData)
		case sig := <-signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		logrus.Errorf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		c <- msg
	}

}

func hearChat(room, ID string, payload *entity.PayloadStateFull) {
	if payload.Room != room {
		return
	}
	if payload.Data.To != ID {
		return
	}
	var fromName string
	if payload.Data.ID == ID {
		fromName = "ME"
	} else {
		fromName = payload.Data.ID
	}
	fmt.Println("--------------- RECEIVE 1 MESSAGES ---------------")
	fmt.Println("From: ", fromName)
	fmt.Println("Messages: ", payload.Data.Chat)
	fmt.Println("###################################################")
}
