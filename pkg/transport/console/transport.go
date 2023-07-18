package console

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/yumemi-inc/go-mailer"
	"github.com/yumemi-inc/go-mailer/pkg/mime"
)

type Transport struct {
	writer io.Writer
}

func NewTransport() *Transport {
	return NewTransportWithWriter(os.Stdout)
}

func NewTransportWithWriter(writer io.Writer) *Transport {
	return &Transport{
		writer: writer,
	}
}

func (t *Transport) Send(email mailer.Email) (*mailer.SentMessage, error) {
	addressesToString := func(addresses []mime.Address) string {
		parts := make([]string, 0, len(addresses))
		for _, a := range addresses {
			if a.Name == "" {
				parts = append(parts, a.Address)
			} else {
				parts = append(parts, fmt.Sprintf("%s <%s>", a.Name, a.Address))
			}
		}

		return strings.Join(parts, ", ")
	}

	contents := make([]string, 0, len(email.Contents)*2)
	for _, c := range email.Contents {
		contents = append(contents, fmt.Sprintf("===== Content-Type: %s =====", c.Type))

		bytes, err := io.ReadAll(c.Body)
		if err != nil {
			return nil, err
		}

		contents = append(contents, string(bytes))
	}

	_, err := fmt.Fprintf(
		t.writer,
		`----- BEGIN EMAIL MESSAGE -----
From: %s
To: %s
Cc: %s
Bcc: %s
Subject: %s

%s
----- END EMAIL MESSAGE -----
`,
		addressesToString(email.From),
		addressesToString(email.To),
		addressesToString(email.Cc),
		addressesToString(email.Bcc),
		email.Subject,
		strings.Join(contents, "\n"),
	)
	if err != nil {
		return nil, err
	}

	return &mailer.SentMessage{
		ID:       "console",
		Original: email,
	}, nil
}
