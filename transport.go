package mailer

type Transport interface {
	Send(email Email) (*SentMessage, error)
}
