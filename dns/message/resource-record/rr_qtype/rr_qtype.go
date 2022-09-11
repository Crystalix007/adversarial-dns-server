package rrqtype

import "fmt"

type RRQType uint8

const (
	A     RRQType = 1
	NS    RRQType = 2
	MD    RRQType = 3
	MF    RRQType = 4
	CNAME RRQType = 5
	SOA   RRQType = 6
	MB    RRQType = 7
	MG    RRQType = 8
	MR    RRQType = 9
	NULL  RRQType = 10
	WKS   RRQType = 11
	PTR   RRQType = 12
	HINFO RRQType = 13
	MINFO RRQType = 14
	MX    RRQType = 15
	TXT   RRQType = 16
	AXFR  RRQType = 252
	MAILB RRQType = 253
	MAILA RRQType = 254
	ALL   RRQType = 255
)

func (rrq RRQType) String() string {
	switch rrq {
	case A:
		return "A"
	case NS:
		return "NS"
	case MD:
		return "MD"
	case MF:
		return "MF"
	case CNAME:
		return "CNAME"
	case SOA:
		return "SOA"
	case MB:
		return "MB"
	case MG:
		return "MG"
	case MR:
		return "MR"
	case NULL:
		return "NULL"
	case WKS:
		return "WKS"
	case PTR:
		return "PTR"
	case HINFO:
		return "HINFO"
	case MINFO:
		return "MINFO"
	case MX:
		return "MX"
	case TXT:
		return "TXT"
	case AXFR:
		return "AXFR"
	case MAILB:
		return "MAILB"
	case MAILA:
		return "MAILA"
	case ALL:
		return "ALL"
	}

	panic(fmt.Sprintf("Unknown RRQType: %d", rrq))
}
