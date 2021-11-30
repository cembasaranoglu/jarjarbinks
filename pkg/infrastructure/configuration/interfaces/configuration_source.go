package interfaces

import "time"

type ConfigurationSource interface {
	GetStringValueByKey(key string) string
	GetIntValueByKey(key string) int
	GetInt64ValueByKey(key string) int64
	GetFloatValueByKey(key string) float64
	GetBooleanValueByKey(key string) bool
	GetTimeValueByKey(key string) time.Time
	GetDurationValueByKey(key string) time.Duration
	GetStringArrayValueByKey(key string) []string
	GetIntArrayValueByKey(key string) []int
	GetValueByKey(key string) interface{}
}

