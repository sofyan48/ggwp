package send

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/sofyan48/ggwp/client/entity"
	"github.com/sofyan48/ggwp/client/utils/kafka"
)

type Send struct {
	Kafka kafka.KafkaInterface
}

func SendHandler() *Send {
	return &Send{
		Kafka: kafka.KafkaHandler(),
	}

}

type SendInterface interface {
	Send(room, id, value string) (*entity.ProducerResponse, error)
}

// Send ..
func (handler *Send) Send(room, id, value string) (*entity.ProducerResponse, error) {
	chat := &entity.PayloadStateFull{}
	chatData := &entity.ChatData{}
	result := &entity.ProducerResponse{}
	chatData.Chat = value
	chatData.ID = id
	ID := uuid.New().String()
	chat.Data = chatData
	chat.Room = room
	chat.ID = ID
	messages, err := json.Marshal(chat)
	if err != nil {
		return result, err
	}
	offset, err := handler.Kafka.Producer("chat", string(messages))
	if err != nil {
		return result, err
	}
	result.ID = ID
	result.Offset = offset
	return result, nil
}
