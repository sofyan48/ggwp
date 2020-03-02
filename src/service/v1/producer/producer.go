package producer

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	httpEntity "github.com/sofyan48/ggwp/src/entity/v1/http"
	entity "github.com/sofyan48/ggwp/src/entity/v1/realtime"
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

// ProducerServiceInterface ...
type ProducerServiceInterface interface {
	SendMessages(payload *httpEntity.ProducerRequest) (int64, error)
}

// SendMessages ...
// @data: string
// topic: string
// return int64, error
func (svc *ProduceService) SendMessages(payload *httpEntity.ProducerRequest) (int64, error) {
	data := &entity.PayloadStateFull{}
	data.Data = payload.Data
	data.Target = payload.Target
	data.ID = uuid.New().String()
	messages, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(messages))
	return 1, nil
	// return svc.Kafka.Producer("test_topic", string(messages))

}
