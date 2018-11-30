package domain

type Message struct {
	ID          string `json:"id"`
	Text        string `json:"text" validate:"required"`
	SenderID    string `json:"sender_id" validate:"required,len=20,alphanum"`
	ReceiverID  string `json:"receiver_id" validate:"required,len=20,alphanum"`
	DateCreated string `json:"date_created"`
}

type MessageStore interface {
	Message(string) (*Message, error)
	Create(*Message) error
	Delete(string) error
	Search(...[2]string) ([]*Message, error)
}
