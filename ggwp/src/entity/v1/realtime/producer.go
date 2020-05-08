package realtime

// PayloadStateFull ...
type PayloadStateFull struct {
	ID   string      `json:"id"`
	Room string      `json:"room"`
	Data interface{} `json:"data"`
}

type ChatData struct {
	ID   string `json:"id"`
	Chat string `json:"chat"`
}
