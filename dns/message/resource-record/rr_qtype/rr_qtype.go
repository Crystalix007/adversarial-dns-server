package rrqtype

import "fmt"

type RRQType uint16

const (
	A          RRQType = 1
	NS         RRQType = 2
	MD         RRQType = 3
	MF         RRQType = 4
	CNAME      RRQType = 5
	SOA        RRQType = 6
	MB         RRQType = 7
	MG         RRQType = 8
	MR         RRQType = 9
	NULL       RRQType = 10
	WKS        RRQType = 11
	PTR        RRQType = 12
	HINFO      RRQType = 13
	MINFO      RRQType = 14
	MX         RRQType = 15
	TXT        RRQType = 16
	RP         RRQType = 17
	AFSDB      RRQType = 18
	SIG        RRQType = 24
	KEY        RRQType = 25
	AAAA       RRQType = 28
	LOC        RRQType = 29
	SRV        RRQType = 33
	NAPTR      RRQType = 35
	KX         RRQType = 36
	CERT       RRQType = 37
	DNAME      RRQType = 39
	OPT        RRQType = 41
	APL        RRQType = 42
	DS         RRQType = 43
	SSHFP      RRQType = 44
	IPSECKEY   RRQType = 45
	RRSIG      RRQType = 46
	NSEC       RRQType = 47
	DNSKEY     RRQType = 48
	DHCID      RRQType = 49
	NSEC3      RRQType = 50
	NSEC3PARAM RRQType = 51
	TLSA       RRQType = 52
	SMIMEA     RRQType = 53
	HIP        RRQType = 55
	CDS        RRQType = 59
	CDNSKEY    RRQType = 60
	OPENPGPKEY RRQType = 61
	CSYNC      RRQType = 62
	ZONEMD     RRQType = 63
	SVCB       RRQType = 64
	HTTPS      RRQType = 65
	EUI48      RRQType = 108
	EUI64      RRQType = 109
	TKEY       RRQType = 249
	TSIG       RRQType = 250
	IFXR       RRQType = 251
	AXFR       RRQType = 252
	MAILB      RRQType = 253
	MAILA      RRQType = 254
	ALL        RRQType = 255
	URI        RRQType = 256
	CAA        RRQType = 257
	TA         RRQType = 32768
	DLV        RRQType = 32769
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
	case IFXR:
		return "IFXR"
	case AXFR:
		return "AXFR"
	case MAILB:
		return "MAILB"
	case MAILA:
		return "MAILA"
	case ALL:
		return "ALL"
	case URI:
		return "URI"
	case CAA:
		return "CAA"
	case TA:
		return "TA"
	case DLV:
		return "DLV"
	}

	panic(fmt.Sprintf("Unknown RRQType: %d", rrq))
}
