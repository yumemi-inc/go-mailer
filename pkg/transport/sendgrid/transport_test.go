package sendgrid

import (
	"bytes"
	"os"
	"testing"

	"github.com/sendgrid/sendgrid-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yumemi-inc/go-mailer"
	"github.com/yumemi-inc/go-mailer/pkg/mime"
)

func TestTransport_Send(t *testing.T) {
	transport := Transport{
		Client: sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY")),
	}

	email := mailer.Email{
		Subject: "Testing SendGrid email transport",
		Sender: mime.Address{
			Address: "me@s6n.jp",
			Name:    "Testing Sender Address",
		},
		From: []mime.Address{
			{
				Address: "me@s6n.jp",
				Name:    "Testing From Address",
			},
		},
		To: []mime.Address{
			{
				Address: "me+testing-to@s6n.jp",
				Name:    "Testing To Address",
			},
		},
		ReplyTo: []mime.Address{
			{
				Address: "me+testing-reply-to@s6n.jp",
				Name:    "Testing Reply-To Address",
			},
		},
		Contents: []mailer.Content{
			{
				Type: mailer.ContentTypeTextPlain,
				Body: bytes.NewBufferString("Hello from go-sender!"),
			},
			{
				Type: mailer.ContentTypeTextHTML,
				Body: bytes.NewBufferString("<h1>Hello from HTML!</h1>"),
			},
		},
	}

	sentMessage, err := transport.Send(email)

	require.NoError(t, err)
	assert.NotNil(t, sentMessage)
}
