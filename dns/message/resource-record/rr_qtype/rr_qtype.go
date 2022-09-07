package rrqtype

type RRQType int8

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
