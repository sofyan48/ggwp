package realtime

// PayloadStateFull ...
type PayloadStateFull struct {
	ID     string      `json:"id"`
	Target string      `json:"target"`
	Data   interface{} `json:"data"`
}
