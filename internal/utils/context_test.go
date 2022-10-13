package utils

import (
	"context"
	"testing"
)

func TestStrGetValueFromCtx(t *testing.T) {
	type ContextKey string

	var key ContextKey = "key"

	type args struct {
		ctx context.Context
		key interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case string",
			args: args{
				ctx: context.WithValue(context.Background(), "key", "value"),
				key: "key",
			},
			want: "value",
		},
		{
			name: "case string (invalid)",
			args: args{
				ctx: context.WithValue(context.Background(), "key", "value"),
				key: key,
			},
			want: "",
		},
		{
			name: "case int",
			args: args{
				ctx: context.WithValue(context.Background(), 123, "value"),
				key: 123,
			},
			want: "value",
		},
		{
			name: "case object",
			args: args{
				ctx: context.WithValue(context.Background(), key, "value"),
				key: key,
			},
			want: "value",
		},
		{
			name: "case object (invalid)",
			args: args{
				ctx: context.WithValue(context.Background(), key, "value"),
				key: "key",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStrValueFromCtx(tt.args.ctx, tt.args.key); got != tt.want {
				t.Errorf("GetStrValueFromCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
