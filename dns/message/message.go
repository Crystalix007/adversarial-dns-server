package message

import (
	"fmt"

	"github.com/Crystalix007/adversarial-dns-server/dns/message/header"
	"github.com/Crystalix007/adversarial-dns-server/dns/message/question"
	resourcerecord "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record"
	"go.uber.org/zap/zapcore"
)

type Message struct {
	Header     header.Header
	Question   question.Questions
	Answer     resourcerecord.RRs
	Authority  resourcerecord.RRs
	Additional resourcerecord.RRs
}

func Decode(b []byte) (*Message, error) {
	h, err := header.Decode(b[:header.HeaderSize])
	if err != nil {
		return nil, err
	}

	var m Message
	m.Header = *h

	q, _, err := question.Decode(b[header.HeaderSize:], m.Header.QDCount.QDCount)
	if err != nil {
		return nil, err
	}

	m.Question = q

	// if len(remainingBytes) != 0 {
	// 	return nil, fmt.Errorf("dns/message: failed to decode message, had %d bytes remaining", len(remainingBytes))
	// }

	return &m, nil
}

func (m Message) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	err := enc.AddObject("Header", m.Header)
	if err != nil {
		return fmt.Errorf("dns/message: failed to encode Header: %w", err)
	}

	err = enc.AddArray("Question", m.Question)

	if err != nil {
		return fmt.Errorf("dns/message: failed to encode Question: %w", err)
	}

	// TODO: Enable these fields.
	// err = enc.AddArray("Answer", m.Answer)

	// if err != nil {
	// 	return fmt.Errorf("dns/message: failed to encode Answer: %w", err)
	// }

	// err = enc.AddArray("Authority", m.Authority)

	// if err != nil {
	// 	return fmt.Errorf("dns/message: failed to encode Authority: %w", err)
	// }

	// err = enc.AddArray("Additional", m.Additional)

	// if err != nil {
	// 	return fmt.Errorf("dns/message: failed to encode Additional: %w", err)
	// }

	return nil
}
