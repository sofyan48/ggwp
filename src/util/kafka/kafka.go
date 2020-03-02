package kafka


type Kafka struct{}

func KafkaHandler() *Kafka {
	return &Kafka{}
}

type KafkaInterface interface{
	Consumer(topic string, data interface{}) (interface, error)
	Producer(topic string, data interface{}) (interface, error)
}

func (kf *Kafka) Consumer(topic string, data interface{}) (interface, error) {

}


func (kf *Kafka) Producer(topic string, data interface{}) (interface, error) {
	
}