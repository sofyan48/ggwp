package entity

// PayloadStateFull ...
type PayloadStateFull struct {
	ID   string    `json:"id"`
	Room string    `json:"room"`
	Data *ChatData `json:"data"`
}

type ChatData struct {
	ID   string `json:"id"`
	To   string `json:"to"`
	Chat string `json:"chat"`
}

// ProducerResponse ...
type ProducerResponse struct {
	Offset int64  `json:"offset"`
	ID     string `json:"ID"`
}
