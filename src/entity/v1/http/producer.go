package http

// ProducerRequest ...
type ProducerRequest struct {
	Target string      `json:"target"`
	Data   interface{} `json:"data"`
}

// ProducerResponse ...
type ProducerResponse struct {
	Offset int64  `json:"offset"`
	ID     string `json:"ID"`
}
