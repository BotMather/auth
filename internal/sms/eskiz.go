package sms

type Eskiz struct{}

func NewEskiz() SmsProvider {
	return &Eskiz{}
}

func (e *Eskiz) Send(phone string, msg string) error {
	return nil
}
