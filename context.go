package fumble

import (
	"context"

	"github.com/mhilmyh/fumble/priority"
	"github.com/mhilmyh/fumble/severity"
)

type Key string

const (
	ContextKey Key = "fumble.ContextKey"
)

type contextValue struct {
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
	return context.WithValue(ctx, ContextKey, contextValue{
		scope:    scope,
		loc:      loc,
		fun:      fun,
		priority: prio,
		severity: sever,
	})
}

func GetContext(ctx context.Context) (contextValue, bool) {
	value, exist := ctx.Value(ContextKey).(contextValue)
	return value, exist
}

func (cv *contextValue) Scope() string {
	return cv.scope
}

func (cv *contextValue) Location() string {
	return cv.loc
}

func (cv *contextValue) Function() string {
	return cv.fun
}

func (cv *contextValue) Priority() Level {
	if cv.priority == nil {
		return priority.None
	}
	return cv.priority
}

func (cv *contextValue) Severity() Level {
	if cv.severity == nil {
		return severity.None
	}
	return cv.severity
}
