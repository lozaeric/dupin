package domain

type Message struct {
	ID         string `json:"id,omitempty"`
	Text       string `json:"text"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Timestamp  string `json:"timestamp"`
}

type MessageStore interface {
	Message(string) (*Message, error)
	CreateMessage(*Message) error
	DeleteMessage(string) error
	Validate(*Message) error
	SearchMessages(...[2]string) ([]*Message, error)
}
