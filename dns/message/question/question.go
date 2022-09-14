package question

import (
	"errors"
	"fmt"

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

func Decode(b []byte, qdCount uint16) (Questions, []byte, error) {
	qs := Questions{}

	for i := uint16(0); i < qdCount; i++ {
		q, remainingBytes, err := decodeQuestion(b)
		if err != nil {
			return nil, nil, err
		}

		qs = append(qs, *q)
		b = remainingBytes
	}

	return qs, b, nil
}

func decodeQuestion(b []byte) (*Question, []byte, error) {
	names, remainingBytes, err := namelabel.Decode(b[:256])
	if err != nil {
		return nil, nil, err
	}

	q := Question{
		QName: names,
	}

	// Append any decode-limited bytes
	if len(b) >= 256 {
		remainingBytes = append(remainingBytes, b[256:]...)
	}

	if len(remainingBytes) < 4 {
		return nil, nil, ErrTruncatedQuestion
	}

	q.QType = rrqtype.RRQType(remainingBytes[0]<<1 + remainingBytes[1])
	q.QClass = qclass.QClass(remainingBytes[2]<<1 + remainingBytes[3])

	return &q, remainingBytes[4:], nil
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
