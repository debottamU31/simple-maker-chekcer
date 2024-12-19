package mailer

import "fmt"

type Mailer interface {
	Send(to, content string) error
}

type consoleMailer struct{}

func NewConsoleMailer() Mailer {
	return &consoleMailer{}
}

func (m *consoleMailer) Send(to, content string) error {
	fmt.Printf("Sending message to %s: %s\n", to, content)
	return nil
}
