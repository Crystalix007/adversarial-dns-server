package resourcerecord

import (
	"errors"
	"fmt"

	"github.com/encodingx/binary"
)

var (
	ErrNoNullRootLabel = errors.New("dns/packet/resource-record: no null root label found")
	ErrMissingData     = errors.New("dns/packet/resource-record: missing bytes required to construct the RR")
	ErrMissingRData    = errors.New("dns/packet/resource-record: missing bytes required to construct the RData field")
)

type RR_Internal struct {
	Type     uint16 `bitfield:"16"`
	Class    uint16 `bitfield:"16"`
	TTL      int32  `bitfield:"32"`
	RDLength uint16 `bitfield:"16"`
}

type Name_Label struct {
	Size  uint8
	Label []byte
}

type RR struct {
	Name []Name_Label
	RR_Internal
	RData []byte
}

func Decode(b []byte) (*RR, error) {
	names, remainingBytes, err := decodeNames(b)
	if err != nil {
		return nil, err
	}

	var rr RR
	rr.Name = names

	err = binary.Unmarshal(remainingBytes, &rr.RR_Internal)

	if err != nil {
		return nil, fmt.Errorf("dns/packet/resource-record: couldn't unmarshal RR: %w", err)
	}

	return &rr, nil
}

func decodeNames(b []byte) ([]Name_Label, []byte, error) {
	if len(b) == 0 {
		return nil, nil, ErrNoNullRootLabel
	}

	l := Name_Label{
		Size:  0,
		Label: []byte{},
	}

	if b[0] == 0 {
		return []Name_Label{l}, b[1:], nil
	}

	l.Size = b[0]
	b = b[1:]

	if int(l.Size) > len(b) {
		return nil, nil, ErrNoNullRootLabel
	}

	l.Label = b[:l.Size]
	b = b[l.Size:]

	labels := []Name_Label{l}

	recursedLabels, remainingBytes, err := decodeNames(b)
	if err != nil {
		return nil, nil, err
	}

	labels = append(labels, recursedLabels...)

	return recursedLabels, remainingBytes, nil
}
