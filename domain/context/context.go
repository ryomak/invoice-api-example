package context

import (
	"context"
)

var (
	authIDKey = "AUTH_ID"
)

func WithAuth(ctx context.Context, auth *Auth) context.Context {
	return context.WithValue(ctx, &authIDKey, auth)
}

func GetAuth(ctx context.Context) *Auth {
	if v := ctx.Value(&authIDKey); v != nil {
		return v.(*Auth)
	}
	return nil
}
