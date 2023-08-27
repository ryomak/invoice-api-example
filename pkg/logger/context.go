package logger

import (
	"context"
)

var (
	contextKey = "USER_CONTEXT"
)

func WithUser(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, &contextKey, userID)
}

func User(ctx context.Context) string {
	if v := ctx.Value(&contextKey); v != nil {
		return v.(string)
	}
	return ""
}
