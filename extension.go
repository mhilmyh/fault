package fault

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
)

type Extension func(e *object)

func WithContext(ctx context.Context, replace bool) Extension {
	return func(e *object) {
		if value, ok := GetContext(ctx); ok {
			if replace || e.scope == "" {
				e.scope = value.scope
			}
			if replace || e.fun == "" {
				e.fun = value.fun
			}
			if replace || e.loc == "" {
				e.loc = value.loc
			}
			if replace || e.priority == nil || e.priority.Empty() {
				e.priority = value.priority
			}
			if replace || e.severity == nil || e.severity.Empty() {
				e.severity = value.severity
			}
		}
	}
}

func WithScope(scope string) Extension {
	return func(e *object) {
		e.scope = scope
	}
}

func WithSeverity(level Level) Extension {
	return func(e *object) {
		e.severity = level
	}
}

func WithPriority(level Level) Extension {
	return func(e *object) {
		e.priority = level
	}
}

func WithParent(parent error) Extension {
	return func(e *object) {
		e.parent = parent
	}
}

func WithCaller(skip int) Extension {
	return func(e *object) {
		e.loc, e.fun = getCaller(skip + 2)
	}
}

func getCaller(skip int) (location string, function string) {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if ok {
		location = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		function = fn.Name()
	}
	return
}
