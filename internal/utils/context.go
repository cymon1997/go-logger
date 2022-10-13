package utils

import "context"

func GetStrValueFromCtx(ctx context.Context, key interface{}) string {
	value, ok := ctx.Value(key).(string)
	if ok {
		return value
	}
	return ""
}
