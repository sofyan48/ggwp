package http

// ProducerRequest ...
type ProducerRequest struct {
	Room string    `json:"room"`
	Data *ChatData `json:"data"`
}

// ProducerResponse ...
type ProducerResponse struct {
	Offset int64  `json:"offset"`
	ID     string `json:"ID"`
}

type ChatData struct {
	ID   string `json:"id"`
	Chat string `json:"chat"`
}
