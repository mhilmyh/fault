package fumble

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/mhilmyh/fumble/priority"
	"github.com/mhilmyh/fumble/severity"
)

type Error interface {
	Error() string
	Message() string
	Scope() string
	Function() string
	Location() string
	Severity() Level
	Priority() Level
	Parent() error
}

type Level interface {
	String() string
	Int() int
	Valid() bool
	Label() string
}

type object struct {
	msg      string
	scope    string
	fun      string
	loc      string
	priority Level
	severity Level
	parent   error
}

func New(ctx context.Context, msg string, extensions ...Extension) Error {
	return Raw(msg, WithContext(ctx), extensions...)
}

func Raw(msg string, extensions ...Extension) Error {
	result := &object{msg: msg}
	for _, extension := range extensions {
		extension(result)
	}
	return result
}

// Error: [scope:function][priority][severity](location) msg...
func (o *object) Error() string {
	var sb strings.Builder

	if len(o.Scope()) > 0 || len(o.Function()) > 0 {
		sb.WriteRune('[')
	}
	if len(o.Scope()) > 0 {
		sb.WriteString(o.Scope())
	}
	if len(o.Scope()) > 0 && len(o.Function()) > 0 {
		sb.WriteRune(':')
	}
	if len(o.Function()) > 0 {
		sb.WriteString(o.Function())
	}
	if len(o.Scope()) > 0 || len(o.Function()) > 0 {
		sb.WriteRune(']')
	}

	if o.Priority().Valid() {
		sb.WriteRune('[')
		sb.WriteString(o.Priority().Label())
		sb.WriteRune(']')
	}

	if o.Severity().Valid() {
		sb.WriteRune('[')
		sb.WriteString(o.Severity().Label())
		sb.WriteRune(']')
	}

	if len(o.Location()) > 0 {
		sb.WriteRune('(')
		sb.WriteString(o.Location())
		sb.WriteRune(')')
	}

	if sb.Len() > 0 {
		sb.WriteRune(' ')
	}
	if len(o.Message()) > 0 {
		sb.WriteString(o.Message())
	}

	return sb.String()
}

func (o *object) Message() string {
	return o.msg
}

func (o *object) Scope() string {
	return o.scope
}

func (o *object) Function() string {
	return o.fun
}

func (o *object) Location() string {
	return o.loc
}

func (o *object) Severity() Level {
	if o.severity == nil {
		return severity.Safe
	}
	return o.severity
}

func (o *object) Priority() Level {
	if o.priority == nil {
		return priority.None
	}
	return o.priority
}

func (o *object) Parent() error {
	return o.parent
}

func (o *object) MarshalJSON() ([]byte, error) {
	var sb strings.Builder

	sb.WriteString("{")

	if scope, hasValue := trimSpace(o.scope); hasValue {
		sb.WriteString(`"scope":`)
		sb.WriteString(strconv.Quote(scope))
	}

	if fun, hasValue := trimSpace(o.fun); hasValue {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(`"function":`)
		sb.WriteString(strconv.Quote(fun))
	}

	if loc, hasValue := trimSpace(o.loc); hasValue {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(`"location":`)
		sb.WriteString(strconv.Quote(loc))
	}

	if msg, hasValue := trimSpace(o.msg); hasValue {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(`"message":`)
		sb.WriteString(strconv.Quote(msg))
	}

	if o.priority != nil {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(`"priority":`)
		sb.WriteString(strconv.Quote(o.priority.String()))
	}

	if o.severity != nil {
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(`"severity":`)
		sb.WriteString(strconv.Quote(o.severity.String()))
	}

	if o.parent != nil {
		parent, _ := json.Marshal(o.parent)
		if sb.Len() > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(`"parent":`)
		sb.Write(parent)
	}

	sb.WriteString("}")

	return []byte(sb.String()), nil
}
