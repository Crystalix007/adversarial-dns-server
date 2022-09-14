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
	TTL uint32 `bitfield:"32"`
}

type RR_RDLength struct {
	RDLength uint16 `bitfield:"16"`
}

type _RR_Internal struct {
	Type     RR_Type     `word:"16"`
	Class    RR_Class    `word:"16"`
	TTL      RR_TTL      `word:"32"`
	RDLength RR_RDLength `word:"16"`
}

type RR_Internal struct {
	Type     rrtype.RRType
	Class    class.Class
	TTL      int32
	RDLength uint16
}

const RR_Internal_Length = 10

type RR struct {
	Name namelabel.NameLabels
	RR_Internal
	RData []byte
}

type RRs []RR

func Decode(b []byte, rrCount uint16) (RRs, []byte, error) {
	rrs := RRs{}

	for i := 0; i < int(rrCount); i++ {
		rr, remainingBytes, err := decodeRR(b)
		if err != nil {
			return nil, nil, err
		}

		rrs = append(rrs, *rr)
		b = remainingBytes
	}

	return rrs, b, nil
}

func decodeRR(b []byte) (*RR, []byte, error) {
	names, remainingBytes, err := namelabel.Decode(b[:256])
	if err != nil {
		return nil, nil, err
	}

	// Append any bytes we limited from the name decoding.
	if len(b) >= 256 {
		remainingBytes = append(remainingBytes, b[256:]...)
	}

	// If we don't have enough bytes to decode the internal data structure,
	// return error.
	if len(remainingBytes) < RR_Internal_Length {
		return nil, nil, ErrMissingData
	}

	rr := RR{
		Name: names,
	}

	var _rrInternal _RR_Internal

	err = binary.Unmarshal(remainingBytes[:RR_Internal_Length], &_rrInternal)

	if err != nil {
		return nil, nil, fmt.Errorf("dns/packet/resource-record: couldn't unmarshal RR: %w", err)
	}

	rr.RR_Internal = _rrInternal.RR_Internal()

	remainingBytes = remainingBytes[RR_Internal_Length:]

	// We don't have enough to decode the RData field.
	if len(remainingBytes) < int(rr.RDLength) {
		return nil, nil, ErrMissingRData
	}

	rr.RData = remainingBytes[:rr.RDLength]
	remainingBytes = remainingBytes[rr.RDLength:]

	return &rr, remainingBytes, nil
}

func (r RR) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	err := enc.AddArray("Name", r.Name)
	if err != nil {
		return fmt.Errorf("dns/message/resource-record: failed to encode Name: %w", err)
	}

	enc.AddString("Type", r.Type.String())
	enc.AddString("Class", r.Class.String())
	enc.AddInt32("TTL", r.TTL)
	enc.AddUint16("RDLength", r.RDLength)

	return err
}

func (rrs RRs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, rr := range rrs {
		err := enc.AppendObject(rr)
		if err != nil {
			return fmt.Errorf("dns/messag/resource-record: failed to encode RR: %w", err)
		}
	}

	return nil
}

func (rr _RR_Internal) RR_Internal() RR_Internal {
	return RR_Internal{
		Type:     rr.Type.Type,
		Class:    rr.Class.Class,
		TTL:      int32(rr.TTL.TTL),
		RDLength: rr.RDLength.RDLength,
	}
}
