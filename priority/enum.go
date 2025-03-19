package priority

import "strconv"

type Enum int

const (
	None Enum = iota
	Low
	Medium
	High
)

func (l Enum) String() string {
	switch l {
	default:
		return "unknown"
	case None:
		return "none"
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	}
}

func (l Enum) Int() int {
	return int(l)
}

func (l Enum) Valid() bool {
	return l >= Low && l <= High
}

func (l Enum) Label() string {
	return "P" + strconv.Itoa(l.Int())
}

func (l Enum) Empty() bool {
	return l == None
}
