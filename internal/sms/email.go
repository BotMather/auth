package sms

type Email struct{}

func NewEmail() SmsProvider {
	return &Email{}
}

func (e *Email) Send(email string, msg string) error {
	return nil
}
