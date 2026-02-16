package sms

type Playmobile struct{}

func NewPlaymobile() SmsProvider {
	return &Playmobile{}
}

func (e *Playmobile) Send(phone string, msg string) error {
	return nil
}
