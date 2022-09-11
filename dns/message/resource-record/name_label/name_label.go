package namelabel

import (
	"errors"

	"go.uber.org/zap/zapcore"
)

var ErrNoNullRootLabel = errors.New("dns/packet/resource-record: no null root label found")

type NameLabel struct {
	Size  uint8
	Label []byte
}

func Decode(b []byte) ([]NameLabel, []byte, error) {
	if len(b) == 0 {
		return nil, nil, ErrNoNullRootLabel
	}

	l := NameLabel{
		Size:  0,
		Label: []byte{},
	}

	if b[0] == 0 {
		return []NameLabel{l}, b[1:], nil
	}

	l.Size = b[0]
	b = b[1:]

	if int(l.Size) > len(b) {
		return nil, nil, ErrNoNullRootLabel
	}

	l.Label = b[:l.Size]
	b = b[l.Size:]

	labels := []NameLabel{l}

	recursedLabels, remainingBytes, err := Decode(b)
	if err != nil {
		return nil, nil, err
	}

	labels = append(labels, recursedLabels...)

	return labels, remainingBytes, nil
}

type NameLabels []NameLabel

func (ns NameLabels) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, n := range ns {
		enc.AppendByteString(n.Label)
	}

	return nil
}
