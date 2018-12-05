package domain

type Message struct {
	ID          string `json:"id" bson:"id"`
	Text        string `json:"text" bson:"text" validate:"required"`
	SenderID    string `json:"sender_id" bson:"sender_id" validate:"required,len=20,alphanum"`
	ReceiverID  string `json:"receiver_id" bson:"receiver_id" validate:"required,len=20,alphanum"`
	DateCreated string `json:"date_created" bson:"date_created"`
}

type MessageStore interface {
	Message(string) (*Message, error)
	Create(*Message) error
	Search(field, value string) ([]*Message, error)
}
