package fault

import (
	"context"

	"github.com/mhilmyh/fault/priority"
	"github.com/mhilmyh/fault/severity"
)

type Key string

const (
	ContextKey Key = "fault.ContextKey"
)

type ContextValue struct {
	scope    string
	loc      string
	fun      string
	priority Level
	severity Level
}

func InitContext(ctx context.Context, scope string, args ...any) context.Context {
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
	return context.WithValue(ctx, ContextKey, ContextValue{
		scope:    scope,
		loc:      loc,
		fun:      fun,
		priority: prio,
		severity: sever,
	})
}

func GetContext(ctx context.Context) (ContextValue, bool) {
	value, exist := ctx.Value(ContextKey).(ContextValue)
	return value, exist
}

func (cv *ContextValue) Scope() string {
	return cv.scope
}

func (cv *ContextValue) Location() string {
	return cv.loc
}

func (cv *ContextValue) Function() string {
	return cv.fun
}

func (cv *ContextValue) Priority() Level {
	if cv.priority == nil {
		return priority.None
	}
	return cv.priority
}

func (cv *ContextValue) Severity() Level {
	if cv.severity == nil {
		return severity.None
	}
	return cv.severity
}
