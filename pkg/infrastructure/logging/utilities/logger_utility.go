package utilities

import (
	"errors"
	"github.com/xiam/to"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jarjarbinks/pkg/infrastructure/logging/enumeration"
	"time"
)

func ToZapFields(parameters ...map[string]interface{}) *[]zap.Field {
	var zapFields []zap.Field
	if parameters != nil && len(parameters) > 0 {
		for _, parameterItem := range parameters {
			for key, value := range parameterItem {
				if field, err := toZapField(key, value); err == nil {
					zapFields = append(zapFields, field)
				}

			}
		}
		return &zapFields
	}
	return &zapFields
}

func toZapField(key string, value interface{}) (zap.Field, error) {
	if len(key) == 0 {
		return zap.String("", ""), errors.New("key cannot be nil")
	}
	switch value.(type) {
	case int:
		return zap.Int(key, to.Int(value)), nil
	case float64:
		return zap.Float64(key, to.Float64(value)), nil
	case string:
		return zap.String(key, to.String(value)), nil
	case time.Time:
		return zap.Time(key, to.Time(value)), nil
	case bool:
		return zap.Bool(key, to.Bool(value)), nil
	default:
		return zap.String(key, to.String(value)), nil
	}
}

func ToZapLogLevel(level string) zapcore.Level {
	switch level {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	case "PANIC":
		return zapcore.PanicLevel
	case "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func ToZapLevelEncoder(encoder enumeration.EncodeLevel) zapcore.LevelEncoder {
	switch encoder {
	case enumeration.Lowercase:
		return zapcore.LowercaseLevelEncoder
	case enumeration.Camelcase:
		return zapcore.CapitalLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func ToZapTimeEncoder(encoder enumeration.EncodeTime) zapcore.TimeEncoder {
	switch encoder {
	case enumeration.RFC3339Nano:
		return zapcore.RFC3339NanoTimeEncoder
	case enumeration.RFC3339:
		return zapcore.RFC3339TimeEncoder
	case enumeration.ISO8601:
		return zapcore.ISO8601TimeEncoder
	case enumeration.Milliseconds:
		return zapcore.EpochMillisTimeEncoder
	case enumeration.Nanoseconds:
		return zapcore.EpochNanosTimeEncoder
	default:
		return zapcore.EpochTimeEncoder
	}
}

func ToZapDurationEncoder(encoder enumeration.EncodeDuration) zapcore.DurationEncoder {
	switch encoder {
	case enumeration.StringDuration:
		return zapcore.StringDurationEncoder
	case enumeration.NanosecondDuration:
		return zapcore.NanosDurationEncoder
	case enumeration.MillisecondDuration:
		return zapcore.NanosDurationEncoder
	default:
		return zapcore.SecondsDurationEncoder
	}
}

func ToZapCallerEncoder(encoder enumeration.EncodeCaller) zapcore.CallerEncoder {
	switch encoder {
	case enumeration.LongestFunctionName:
		return zapcore.FullCallerEncoder
	case enumeration.ShortestFunctionName:
		return zapcore.ShortCallerEncoder
	default:
		return zapcore.FullCallerEncoder
	}
}
