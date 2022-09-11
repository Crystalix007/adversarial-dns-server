package header

import (
	"fmt"

	"github.com/Crystalix007/adversarial-dns-server/dns/message/header/opcode"
	"github.com/Crystalix007/adversarial-dns-server/dns/message/header/rcode"
	"github.com/encodingx/binary"
	"go.uber.org/zap/zapcore"
)

type Header_ID struct {
	ID uint16 `bitfield:"16"`
}

type Header_Flags struct {
	QR     bool  `bitfield:"1"`
	Opcode uint8 `bitfield:"4"`
	AA     bool  `bitfield:"1"`
	TC     bool  `bitfield:"1"`
	RD     bool  `bitfield:"1"`
	RA     bool  `bitfield:"1"`
	Z      uint8 `bitfield:"3"`
	RCode  uint8 `bitfield:"4"`
}

type Header_QDCount struct {
	QDCount uint16 `bitfield:"16"`
}

type Header_ANCount struct {
	ANCount uint16 `bitfield:"16"`
}

type Header_NSCount struct {
	NSCount uint16 `bitfield:"16"`
}

type Header_ARCount struct {
	ARCount uint16 `bitfield:"16"`
}

type Header struct {
	ID      Header_ID      `word:"16"`
	Flags   Header_Flags   `word:"16"`
	QDCount Header_QDCount `word:"16"`
	ANCount Header_ANCount `word:"16"`
	NSCount Header_NSCount `word:"16"`
	ARCount Header_ARCount `word:"16"`
}

const HeaderSize = 12

func Decode(b []byte) (*Header, error) {
	var h Header
	err := binary.Unmarshal(b, &h)
	if err != nil {
		return nil, fmt.Errorf("dns/message/header: failed to decode header: %w", err)
	}

	return &h, nil
}

func (h Header_Flags) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddBool("QR", h.QR)
	enc.AddString("Opcode", opcode.Opcode(h.Opcode).String())
	enc.AddBool("AA", h.AA)
	enc.AddBool("TC", h.TC)
	enc.AddBool("RD", h.RD)
	enc.AddBool("RA", h.RA)
	enc.AddUint8("Z", h.Z)
	enc.AddString("RCode", rcode.RCode(h.RCode).String())

	return nil
}

func (h Header) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddUint16("ID", h.ID.ID)

	err := enc.AddObject("Flags", h.Flags)
	if err != nil {
		return fmt.Errorf("dns/message/header: couldn't marshal Flags object: %w", err)
	}

	enc.AddUint16("QDCount", h.QDCount.QDCount)
	enc.AddUint16("ANCount", h.ANCount.ANCount)
	enc.AddUint16("NSCount", h.NSCount.NSCount)
	enc.AddUint16("ARCount", h.ARCount.ARCount)

	return nil
}
