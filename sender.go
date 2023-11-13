package mailer

import (
	"github.com/yumemi-inc/go-mailer/pkg/mime"
)

type Sender struct {
	transport            Transport
	defaultSenderAddress mime.Address
	defaultFromAddress   mime.Address
}

type Option func(*Sender)

func WithDefaultSenderAddress(address mime.Address) Option {
	return func(sender *Sender) {
		sender.defaultSenderAddress = address
	}
}

func WithDefaultFromAddress(address mime.Address) Option {
	return func(sender *Sender) {
		sender.defaultFromAddress = address
	}
}

func NewSender(transport Transport, opts ...Option) *Sender {
	sender := &Sender{
		transport: transport,
	}

	for _, opt := range opts {
		opt(sender)
	}

	return sender
}

func (s *Sender) Send(email Email) (*SentMessage, error) {
	if email.Sender == (mime.Address{}) {
		email.Sender = s.defaultSenderAddress
	}

	if len(email.From) == 0 {
		email.From = []mime.Address{s.defaultFromAddress}
	}

	return s.transport.Send(email)
}
