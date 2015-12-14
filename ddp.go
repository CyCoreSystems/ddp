package ddp

// core structures

type Message map[string]interface{}

func NewMessage(t string, id string) Message {
	msg := make(Message)
	msg["msg"] = t
	msg["id"] = id
	return msg
}

func (m Message) ID() string {
	s, _ := m["id"].(string)
	return s
}

func (m Message) Type() string {
	s, _ := m["msg"].(string)
	return s
}

type Error struct {
	Error   string `json:"error"`
	Reason  string `json:"reason"`
	Details string `json:"details"`

	OffendingMessage string `json:"offendingMessage"`
}
