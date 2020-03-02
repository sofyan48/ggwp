package http

// ProducerRequest ...
type ProducerRequest struct {
	Target string      `json:"target"`
	Data   interface{} `json:"data"`
}
