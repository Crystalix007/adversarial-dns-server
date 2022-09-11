package qclass

type QClass uint8

const (
	IN  QClass = 1
	CS  QClass = 2
	CH  QClass = 3
	HS  QClass = 4
	ALL QClass = 255
)

func (q QClass) String() string {
	switch q {
	case IN:
		return "Internet"
	case CS:
		return "CSNET"
	case CH:
		return "CHAOS"
	case HS:
		return "Hesiod"
	case ALL:
		return "All"
	}

	panic("Unexpected QClass")
}
