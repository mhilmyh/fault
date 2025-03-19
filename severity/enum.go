package severity

import "strconv"

type Enum int

const (
	None Enum = iota
	Safe
	Trivial
	Minor
	Major
)

func (l Enum) String() string {
	switch l {
	default:
		return "unknown"
	case None:
		return "none"
	case Safe:
		return "safe"
	case Trivial:
		return "trivial"
	case Minor:
		return "minor"
	case Major:
		return "major"
	}
}

func (l Enum) Int() int {
	return int(l)
}

func (l Enum) Valid() bool {
	return l >= Safe && l <= Major
}

func (l Enum) Label() string {
	return "S" + strconv.Itoa(l.Int())
}

func (l Enum) Empty() bool {
	return l == None
}
