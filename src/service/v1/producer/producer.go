package producer

import (
	"github.com/sofyan48/ggwp/src/util/kafka"
)

// ProduceService ...
type ProduceService struct {
	Kafka kafka.KafkaInterface
}

// ProducerServiceHAndler ...
func ProducerServiceHAndler() *ProduceService {
	return &ProduceService{
		Kafka: kafka.KafkaHandler(),
	}
}

// ProducerInterface ...
type ProducerInterface interface {
	SendMessages(data, topic string) (int64, error)
}

// SendMessages ...
// @data: string
// topic: string
// return int64, error
func (svc *ProduceService) SendMessages(data, topic string) (int64, error) {
	return svc.Kafka.Producer("test_topic", data)
}
