package class

type Class uint8

const (
	IN Class = 1
	CS Class = 2
	CH Class = 3
	HS Class = 4
)

func (c Class) String() string {
	switch c {
	case IN:
		return "Internet"
	case CS:
		return "CSNET"
	case CH:
		return "CHAOS"
	case HS:
		return "Hesiod"
	}

	panic("Unexpected Class")
}
