package opcode

type Opcode uint8

const (
	Query  Opcode = 0
	IQuery Opcode = 1
	Status Opcode = 2
)

func (o Opcode) String() string {
	switch o {
	case Query:
		return "query"
	case IQuery:
		return "iquery"
	case Status:
		return "status"
	}

	panic("Unknown opcode")
}
