package domain

type Message struct {
	ID          string `json:"id"`
	Text        string `json:"text" validate:"required"`
	SenderID    string `json:"sender_id" validate:"required,len=20"`
	ReceiverID  string `json:"receiver_id" validate:"required,len=20"`
	DateCreated string `json:"date_created"`
}

type MessageStore interface {
	Message(string) (*Message, error)
	CreateMessage(*Message) error
	DeleteMessage(string) error
	SearchMessages(...[2]string) ([]*Message, error)
}
