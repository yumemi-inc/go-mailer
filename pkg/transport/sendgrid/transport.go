package sendgrid

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/yumemi-inc/go-mailer"
)

type Transport struct {
	Client *sendgrid.Client
}

func NewTransport(apiKey string) *Transport {
	return &Transport{
		Client: sendgrid.NewSendClient(apiKey),
	}
}

func (t *Transport) Send(email mailer.Email) (*mailer.SentMessage, error) {
	var replyTo *mail.Email
	if len(email.ReplyTo) > 0 {
		replyTo = &mail.Email{
			Name:    email.ReplyTo[0].Name,
			Address: email.ReplyTo[0].Address,
		}
	}

	contents := make([]*mail.Content, 0)
	for _, c := range email.Contents {
		buf, err := io.ReadAll(c.Body)
		if err != nil {
			return nil, err
		}

		contents = append(contents, mail.NewContent(c.Type, string(buf)))
	}

	var personalizations []*mail.Personalization
	if len(email.To) > 0 || len(email.Cc) > 0 || len(email.Bcc) > 0 {
		personalization := &mail.Personalization{}
		if len(email.To) > 0 {
			to := make([]*mail.Email, 0, len(email.To))
			for _, a := range email.To {
				to = append(to, mail.NewEmail(a.Name, a.Address))
			}

			personalization.To = to
		}

		personalizations = []*mail.Personalization{personalization}
	}

	response, err := t.Client.Send(
		&mail.SGMailV3{
			From:             mail.NewEmail(email.Sender.Name, email.Sender.Address),
			Subject:          email.Subject,
			Personalizations: personalizations,
			Content:          contents,
			ReplyTo:          replyTo,
		},
	)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusAccepted {
		errorResponse := new(ErrorResponse)
		if err := json.Unmarshal([]byte(response.Body), errorResponse); err != nil {
			return nil, err
		}

		return nil, errorResponse
	}

	return &mailer.SentMessage{
		ID:       response.Headers["X-Message-Id"][0],
		Original: email,
	}, nil
}
