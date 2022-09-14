package class

import "fmt"

type Class uint16

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

	return fmt.Sprintf("Unknown (%d)", c)
	// panic("Unexpected Class")
}
