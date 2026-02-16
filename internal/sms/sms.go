package sms

type SmsProvider interface {
	Send(string, string) error
}
