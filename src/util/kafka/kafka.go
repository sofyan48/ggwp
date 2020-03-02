package kafka

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// Kafka ...
type Kafka struct {
}

// KafkaHandler ...
func KafkaHandler() *Kafka {
	return &Kafka{}
}

// KafkaInterface ...
type KafkaInterface interface {
	Consumer(topics []string, signals chan os.Signal) (interface{}, error)
	Producer(topic string, data interface{}) (interface{}, error)
}

// Consumer ...
func (kf *Kafka) Consumer(topics []string, signals chan os.Signal) {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)
	var consumer sarama.Consumer
	for _, topic := range topics {
		partitionList, err := consumer.Partitions(topic)
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
			continue
		}
		for _, partition := range partitionList {
			fmt.Println(partition)
			go consumeMessage(consumer, topic, partition, chanMessage)
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

// Producer ...
func (kf *Kafka) Producer(topic string, msg string) (interface{}, error) {
	var producer sarama.SyncProducer
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	_, _, err := producer.SendMessage(kafkaMsg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
