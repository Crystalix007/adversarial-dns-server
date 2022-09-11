package rrtype

type RRType uint8

const (
	A     RRType = 1
	NS    RRType = 2
	MD    RRType = 3
	MF    RRType = 4
	CNAME RRType = 5
	SOA   RRType = 6
	MB    RRType = 7
	MG    RRType = 8
	MR    RRType = 9
	NULL  RRType = 10
	WKS   RRType = 11
	PTR   RRType = 12
	HINFO RRType = 13
	MINFO RRType = 14
	MX    RRType = 15
	TXT   RRType = 16
)

func (rrt RRType) String() string {
	switch rrt {
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
	}

	panic("Unknown RRType")
}
