package message

import (
	"fmt"

	"github.com/Crystalix007/adversarial-dns-server/buffer"
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

func Decode(b buffer.Buffer) (*Message, error) {
	h, err := header.Decode(&b)
	if err != nil {
		return nil, err
	}

	var m Message
	m.Header = *h

	q, err := question.Decode(&b, m.Header.QDCount.QDCount)
	if err != nil {
		return nil, err
	}

	m.Question = q

	rrs, err := resourcerecord.Decode(&b, m.Header.ANCount.ANCount)
	if err != nil {
		return nil, err
	}

	m.Answer = rrs

	rrs, err = resourcerecord.Decode(&b, m.Header.NSCount.NSCount)
	if err != nil {
		return nil, err
	}

	m.Authority = rrs

	rrs, err = resourcerecord.Decode(&b, m.Header.ARCount.ARCount)
	if err != nil {
		return nil, err
	}

	m.Additional = rrs

	if b.Length() != 0 {
		return nil, fmt.Errorf("dns/message: failed to decode message, had %d bytes remaining", b.Length())
	}

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
	err = enc.AddArray("Answer", m.Answer)

	if err != nil {
		return fmt.Errorf("dns/message: failed to encode Answer: %w", err)
	}

	err = enc.AddArray("Authority", m.Authority)

	if err != nil {
		return fmt.Errorf("dns/message: failed to encode Authority: %w", err)
	}

	err = enc.AddArray("Additional", m.Additional)

	if err != nil {
		return fmt.Errorf("dns/message: failed to encode Additional: %w", err)
	}

	return nil
}
