package message

import (
	"github.com/Crystalix007/adversarial-dns-server/dns/message/header"
	resourcerecord "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record"
	"go.uber.org/zap/zapcore"
)

type (
	Question string
)

type Message struct {
	Header     header.Header
	Question   Question
	Answer     []resourcerecord.RR
	Authority  []resourcerecord.RR
	Additional []resourcerecord.RR
}

func Decode(b []byte) (*Message, error) {
	h, err := header.Decode(b[:header.HeaderSize])
	if err != nil {
		return nil, err
	}

	var m Message
	m.Header = *h

	return &m, nil
}

func (m *Message) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	err := enc.AddObject("Header", m.Header)
	if err != nil {
		return err
	}

	// TODO: Add other portions of message
	return nil
}
