package cache

import (
	"context"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	SetAndGet()
}

func TestHSetAndHGet(t *testing.T) {
	HSetAndHGet()
}

func TestHSet(t *testing.T) {
	type args struct {
		ctx    context.Context
		key    string
		values []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"hset", args{ctx: ctx, key: "hkey", values: []interface{}{"key1", "v1"}}, false},
		{"hset", args{ctx: ctx, key: "hkey", values: []interface{}{"key2", "v2"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := HSet(tt.args.ctx, tt.args.key, tt.args.values...); (err != nil) != tt.wantErr {
				t.Errorf("HSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHGetResult(t *testing.T) {
	type args struct {
		ctx   context.Context
		key   string
		field string
	}
	var tests = []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{ctx: ctx, key: "hkey", field: "key1"}, "v1", false},
		{"", args{ctx: ctx, key: "hkey", field: "key2"}, "v2", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HGetResult(tt.args.ctx, tt.args.key, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("HGetResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HGetResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}
