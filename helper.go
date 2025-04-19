package fault

import (
	"context"
	"strings"
)

func Catch(ctx context.Context, fn func(ctx context.Context) error) (err Error) {
	defer func() {
		if r := recover(); r != nil {
			switch value := r.(type) {
			case string:
				err = New(ctx, value)
			case error:
				err = New(ctx, value.Error(), WithParent(err))
			}
		}
	}()
	e := fn(ctx)
	if e != nil {
		return New(ctx, e.Error(), WithParent(e))
	}
	return nil
}

func trimSpace(raw string) (content string, hasValue bool) {
	s := strings.TrimSpace(raw)
	return s, len(s) > 0
}
