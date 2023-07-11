package mailer

import (
	"io"

	"github.com/yumemi-inc/go-mailer/pkg/mime"
)

const (
	ContentTypeTextPlain = "text/plain"
	ContentTypeTextHTML  = "text/html"
)

type Content struct {
	Type string
	Body io.Reader
}

type Email struct {
	Subject    string
	Sender     mime.Address
	ReturnPath []mime.Address
	From       []mime.Address
	To         []mime.Address
	Cc         []mime.Address
	Bcc        []mime.Address
	ReplyTo    []mime.Address
	Contents   []Content
}
