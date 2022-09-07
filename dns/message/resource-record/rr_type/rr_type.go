package rrtype

type RRType int8

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
