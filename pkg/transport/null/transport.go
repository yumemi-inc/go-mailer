package null

import (
	"github.com/yumemi-inc/go-mailer"
)

type Transport struct{}

func (t *Transport) Send(email mailer.Email) (*mailer.SentMessage, error) {
	return &mailer.SentMessage{
		ID:       "null",
		Original: email,
	}, nil
}
