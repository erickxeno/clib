package cconfig

import (
	"context"
)

// Configuration interface
// Obtain the configuration item by specifying a Key, including a method with specified default value
type Configuration interface {
	GetInt(ctx context.Context, key string) (int, error)
	GetInt32(ctx context.Context, key string) (int32, error)
	GetInt64(ctx context.Context, key string) (int64, error)
	GetBool(ctx context.Context, key string) (bool, error)
	GetFloat64(ctx context.Context, key string) (float64, error)
	GetString(ctx context.Context, key string) (string, error)
	GetStrings(ctx context.Context, key string) ([]string, error)
	GetJson(ctx context.Context, key string, val interface{}) error // val should be ptr

	GetIntWithDefault(ctx context.Context, key string, v int) int
	GetInt32WithDefault(ctx context.Context, key string, v int32) int32
	GetInt64WithDefault(ctx context.Context, key string, v int64) int64
	GetBoolWithDefault(ctx context.Context, key string, v bool) bool
	GetFloat64WithDefault(ctx context.Context, key string, v float64) float64
	GetStringWithDefault(ctx context.Context, key string, v string) string
	GetStringsWithDefault(ctx context.Context, key string, v []string) []string
}
