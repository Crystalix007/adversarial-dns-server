package rcode

type RCode uint8

const (
	NO_ERR      RCode = 0
	FORMAT_ERR  RCode = 1
	SRV_FAILURE RCode = 2
	NAME_ERR    RCode = 3
	NOT_IMPL    RCode = 4
	REFUSED     RCode = 5
)

func (r RCode) String() string {
	switch r {
	case NO_ERR:
		return "no error"
	case FORMAT_ERR:
		return "format error"
	case SRV_FAILURE:
		return "server failure"
	case NAME_ERR:
		return "name error"
	case NOT_IMPL:
		return "not implemented"
	case REFUSED:
		return "refused"
	}

	panic("Unknown RCode")
}
