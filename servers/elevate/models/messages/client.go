package messages

type ClientMsg struct {
	Type	string `json:"type"`
	Payload interface{} `json:"payload"`
}