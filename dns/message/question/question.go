package question

import (
	"errors"
	"fmt"

	"github.com/Crystalix007/adversarial-dns-server/buffer"
	namelabel "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record/name_label"
	"github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record/qclass"
	rrqtype "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record/rr_qtype"
	"go.uber.org/zap/zapcore"
)

var ErrTruncatedQuestion = errors.New("dns/message/question: truncated question data")

type Question struct {
	QName  namelabel.NameLabels
	QType  rrqtype.RRQType
	QClass qclass.QClass
}

type Questions []Question

func Decode(b *buffer.Buffer, qdCount uint16) (Questions, error) {
	qs := Questions{}

	for i := uint16(0); i < qdCount; i++ {
		q, err := decodeQuestion(b)
		if err != nil {
			return nil, err
		}

		qs = append(qs, *q)
	}

	return qs, nil
}

func decodeQuestion(b *buffer.Buffer) (*Question, error) {
	names, remainingBytes, err := namelabel.Decode(b.Dequeue(256))
	if err != nil {
		return nil, err
	}

	q := Question{
		QName: names,
	}

	// Append any decode-limited bytes
	b.Prepend(remainingBytes)

	if b.Length() < 4 {
		return nil, ErrTruncatedQuestion
	}

	qtypeBytes := b.Dequeue(2)
	q.QType = rrqtype.RRQType(qtypeBytes[0]<<1 + qtypeBytes[1])
	qclassBytes := b.Dequeue(2)
	q.QClass = qclass.QClass(qclassBytes[0]<<1 + qclassBytes[1])

	return &q, nil
}

func (q Question) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	err := enc.AddArray("QName", q.QName)
	if err != nil {
		return fmt.Errorf("dns/message/question: failed to encode QName: %w", err)
	}

	enc.AddString("QType", q.QType.String())
	enc.AddString("QClass", q.QClass.String())

	return nil
}

func (qs Questions) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, q := range qs {
		err := enc.AppendObject(q)
		if err != nil {
			return err
		}
	}

	return nil
}
