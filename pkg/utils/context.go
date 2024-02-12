package utils

import "context"

func ExtractDataFromContext(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
