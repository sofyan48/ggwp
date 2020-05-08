package producer

import (
	"encoding/json"

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
	SendMessages(payload *httpEntity.ProducerRequest) (*httpEntity.ProducerResponse, error)
}

// SendMessages ...
// @data: string
// topic: string
// return int64, error
func (svc *ProduceService) SendMessages(payload *httpEntity.ProducerRequest) (*httpEntity.ProducerResponse, error) {
	data := &entity.PayloadStateFull{}
	result := &httpEntity.ProducerResponse{}
	ID := uuid.New().String()
	data.Data = payload.Data
	data.Room = payload.Room
	data.ID = ID
	messages, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	offset, err := svc.Kafka.Producer("chat", string(messages))
	if err != nil {
		return result, err
	}
	result.ID = ID
	result.Offset = offset
	return result, nil

}
