package ctxutil

import "context"

type localeKeyType struct{}

var localeKey = localeKeyType{}

type LocaleContext struct {
	Locale string
	Region string
}

func WithLocale(ctx context.Context, l LocaleContext) context.Context {
	return context.WithValue(ctx, localeKey, l)
}

func Locale(ctx context.Context) (LocaleContext, bool) {
	v, ok := ctx.Value(localeKey).(LocaleContext)
	return v, ok
}
