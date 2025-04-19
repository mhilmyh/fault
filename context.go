package fault

import (
	"context"

	"github.com/mhilmyh/fault/priority"
	"github.com/mhilmyh/fault/severity"
)

type contextKey string

const (
	Key contextKey = "fault.ContextKey"
)

type Value struct {
	scope    string
	loc      string
	fun      string
	priority Level
	severity Level
}

func Init(ctx context.Context, scope string, args ...any) context.Context {
	var (
		prio  Level = priority.None
		sever Level = severity.None
	)
	for _, arg := range args {
		switch x := arg.(type) {
		case priority.Enum:
			prio = x
		case severity.Enum:
			sever = x
		}
	}
	loc, fun := getCaller(1)
	return context.WithValue(ctx, Key, Value{
		scope:    scope,
		loc:      loc,
		fun:      fun,
		priority: prio,
		severity: sever,
	})
}

func GetContext(ctx context.Context) (Value, bool) {
	value, exist := ctx.Value(Key).(Value)
	return value, exist
}

func (cv *Value) Scope() string {
	return cv.scope
}

func (cv *Value) Location() string {
	return cv.loc
}

func (cv *Value) Function() string {
	return cv.fun
}

func (cv *Value) Priority() Level {
	if cv.priority == nil {
		return priority.None
	}
	return cv.priority
}

func (cv *Value) Severity() Level {
	if cv.severity == nil {
		return severity.None
	}
	return cv.severity
}
