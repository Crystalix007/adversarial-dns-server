package rrtype

import "fmt"

type RRType uint16

const (
	A          RRType = 1
	NS         RRType = 2
	MD         RRType = 3
	MF         RRType = 4
	CNAME      RRType = 5
	SOA        RRType = 6
	MB         RRType = 7
	MG         RRType = 8
	MR         RRType = 9
	NULL       RRType = 10
	WKS        RRType = 11
	PTR        RRType = 12
	HINFO      RRType = 13
	MINFO      RRType = 14
	MX         RRType = 15
	TXT        RRType = 16
	RP         RRType = 17
	AFSDB      RRType = 18
	SIG        RRType = 24
	KEY        RRType = 25
	AAAA       RRType = 28
	LOC        RRType = 29
	SRV        RRType = 33
	NAPTR      RRType = 35
	KX         RRType = 36
	CERT       RRType = 37
	DNAME      RRType = 39
	OPT        RRType = 41
	APL        RRType = 42
	DS         RRType = 43
	SSHFP      RRType = 44
	IPSECKEY   RRType = 45
	RRSIG      RRType = 46
	NSEC       RRType = 47
	DNSKEY     RRType = 48
	DHCID      RRType = 49
	NSEC3      RRType = 50
	NSEC3PARAM RRType = 51
	TLSA       RRType = 52
	SMIMEA     RRType = 53
	HIP        RRType = 55
	CDS        RRType = 59
	CDNSKEY    RRType = 60
	OPENPGPKEY RRType = 61
	CSYNC      RRType = 62
	ZONEMD     RRType = 63
	SVCB       RRType = 64
	HTTPS      RRType = 65
	EUI48      RRType = 108
	EUI64      RRType = 109
	TKEY       RRType = 249
	TSIG       RRType = 250
	URI        RRType = 256
	CAA        RRType = 257
	TA         RRType = 32768
	DLV        RRType = 32769
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
	case RP:
		return "RP"
	case AFSDB:
		return "AFSDB"
	case SIG:
		return "SIG"
	case KEY:
		return "KEY"
	case AAAA:
		return "AAAA"
	case LOC:
		return "LOC"
	case SRV:
		return "SRV"
	case NAPTR:
		return "NAPTR"
	case KX:
		return "KX"
	case CERT:
		return "CERT"
	case DNAME:
		return "DNAME"
	case OPT:
		return "OPT"
	case APL:
		return "APL"
	case DS:
		return "DS"
	case SSHFP:
		return "SSHFP"
	case IPSECKEY:
		return "IPSECKEY"
	case RRSIG:
		return "RRSIG"
	case NSEC:
		return "NSEC"
	case DNSKEY:
		return "DNSKEY"
	case DHCID:
		return "DHCID"
	case NSEC3:
		return "NSEC3"
	case NSEC3PARAM:
		return "NSEC3PARAM"
	case TLSA:
		return "TLSA"
	case SMIMEA:
		return "SMIMEA"
	case HIP:
		return "HIP"
	case CDS:
		return "CDS"
	case CDNSKEY:
		return "CDNSKEY"
	case OPENPGPKEY:
		return "OPENPGPKEY"
	case CSYNC:
		return "CSYNC"
	case ZONEMD:
		return "ZONEMD"
	case SVCB:
		return "SVCB"
	case HTTPS:
		return "HTTPS"
	case EUI48:
		return "EUI48"
	case EUI64:
		return "EUI64"
	case TKEY:
		return "TKEY"
	case TSIG:
		return "TSIG"
	case URI:
		return "URI"
	case CAA:
		return "CAA"
	case TA:
		return "TA"
	case DLV:
		return "DLV"
	}

	return fmt.Sprintf("Unknown (%d)", rrt)
}
