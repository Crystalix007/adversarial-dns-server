package resourcerecord

import (
	"errors"
	"fmt"

	class "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record/class"
	namelabel "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record/name_label"
	rrtype "github.com/Crystalix007/adversarial-dns-server/dns/message/resource-record/rr_type"
	"github.com/encodingx/binary"
	"go.uber.org/zap/zapcore"
)

var (
	ErrMissingData  = errors.New("dns/packet/resource-record: missing bytes required to construct the RR")
	ErrMissingRData = errors.New("dns/packet/resource-record: missing bytes required to construct the RData field")
)

type RR_Type struct {
	Type rrtype.RRType `bitfield:"16"`
}

type RR_Class struct {
	Class class.Class `bitfield:"16"`
}

type RR_TTL struct {
	TTL int32 `bitfield:"32"`
}

type RR_RDLength struct {
	RDLength uint16 `bitfield:"16"`
}

type RR_Internal struct {
	Type     RR_Type     `word:"16"`
	Class    RR_Class    `word:"16"`
	TTL      RR_TTL      `word:"32"`
	RDLength RR_RDLength `word:"16"`
}

type RR struct {
	Name namelabel.NameLabels
	RR_Internal
	RData []byte
}

func Decode(b []byte) (*RR, error) {
	names, remainingBytes, err := namelabel.Decode(b[:256])
	if err != nil {
		return nil, err
	}

	// Append any bytes we limited from the name decoding.
	if len(b) >= 256 {
		remainingBytes = append(remainingBytes, b[256:]...)
	}

	rr := RR{
		Name: names,
	}

	err = binary.Unmarshal(remainingBytes[:10], &rr.RR_Internal)

	if err != nil {
		return nil, fmt.Errorf("dns/packet/resource-record: couldn't unmarshal RR: %w", err)
	}

	// TODO: Validate data length.
	rr.RData = remainingBytes[:10]

	return &rr, nil
}

func (r RR) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	err := enc.AddArray("Name", r.Name)
	if err != nil {
		return fmt.Errorf("dns/message/resource-record: failed to encode Name: %w", err)
	}

	enc.AddString("Type", r.Type.Type.String())
	enc.AddString("Class", r.Class.Class.String())
	enc.AddInt32("TTL", r.TTL.TTL)
	enc.AddUint16("RDLength", r.RDLength.RDLength)

	return err
}

type RRs []RR

func (rrs RRs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, rr := range rrs {
		err := enc.AppendObject(rr)
		if err != nil {
			return fmt.Errorf("dns/messag/resource-record: failed to encode RR: %w", err)
		}
	}

	return nil
}
