package send

import (
	"github.com/sofyan48/ggwp/client/utils/kafka"
)

type Send struct {
	Producer kafka.KafkaInterface
}

func SendHandler() *Send {
	return &Send{
		Producer: kafka.KafkaHandler(),
	}

}

type SendInterface interface{}

func (handler *Send) Send() {

}
