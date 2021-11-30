package options

import (
	"jarjarbinks/pkg/infrastructure/logging/enumeration"
)

var (
	defaultEncoding         = "json"
	defaultOutputPaths      = []string{"stdout"}
	defaultErrorOutputPaths = []string{"stdout"}
	defaultLoggingLevel     = "DEBUG"
	defaultTimeKey          = "timestamp"
	defaultLevelKey         = "severity"
	defaultNameKey          = "default-logger"
	defaultCallerKey        = "caller"
	defaultMessageKey       = "message"
	defaultStackTraceKey    = "stacktrace"
	defaultFunctionKey      = "method"
	defaultEncodeLevel      = enumeration.Lowercase
	defaultEncodeTime       = enumeration.RFC3339
	defaultEncodeDuration   = enumeration.MillisecondDuration
	defaultEncodeCaller     = enumeration.ShortestFunctionName
)

type LoggerOptions struct {
	DefaultParameters map[string]interface{}
	Development       bool
	Encoding          string
	OutputPaths       []string
	ErrorOutputPaths  []string
	Level             string
	TimeKey           string
	LevelKey          string
	NameKey           string
	CallerKey         string
	StackTraceKey     string
	MessageKey        string
	FunctionKey       string
	EncodeLevel       *enumeration.EncodeLevel
	EncodeTime        *enumeration.EncodeTime
	EncodeDuration    *enumeration.EncodeDuration
	EncodeCaller      *enumeration.EncodeCaller
}

func New() *LoggerOptions {
	return &LoggerOptions{
		DefaultParameters: nil,
		Development:       false,
		Encoding:          defaultEncoding,
		OutputPaths:       defaultOutputPaths,
		ErrorOutputPaths:  defaultErrorOutputPaths,
		Level:             defaultLoggingLevel,
		TimeKey:           defaultTimeKey,
		LevelKey:          defaultLevelKey,
		NameKey:           defaultNameKey,
		CallerKey:         defaultCallerKey,
		StackTraceKey:     defaultStackTraceKey,
		MessageKey:        defaultMessageKey,
		EncodeLevel:       &defaultEncodeLevel,
		EncodeTime:        &defaultEncodeTime,
		EncodeDuration:    &defaultEncodeDuration,
		EncodeCaller:      &defaultEncodeCaller,
	}
}

func (c *LoggerOptions) UseDefaultEncodingIfNotSpecified() *LoggerOptions {
	if len(c.Encoding) == 0 {
		c.Encoding = defaultEncoding
	}
	return c
}

func (c *LoggerOptions) UseDefaultOutputPathsIfNotSpecified() *LoggerOptions {
	if c.OutputPaths == nil || len(c.OutputPaths) == 0 {
		c.OutputPaths = defaultOutputPaths
	}
	return c
}

func (c *LoggerOptions) UseDefaultErrorOutputPathsIfNotSpecified() *LoggerOptions {
	if c.ErrorOutputPaths == nil || len(c.ErrorOutputPaths) == 0 {
		c.ErrorOutputPaths = defaultErrorOutputPaths
	}
	return c
}

func (c *LoggerOptions) UseDefaultLoggingLevelIfNotSpecified() *LoggerOptions {
	if len(c.Level) == 0 {
		c.Level = defaultLoggingLevel
	}
	return c
}

func (c *LoggerOptions) UseDefaultTimeKeyIfNotSpecified() *LoggerOptions {
	if len(c.TimeKey) == 0 {
		c.TimeKey = defaultTimeKey
	}
	return c
}

func (c *LoggerOptions) UseDefaultLevelKeyIfNotSpecified() *LoggerOptions {
	if len(c.LevelKey) == 0 {
		c.LevelKey = defaultLevelKey
	}
	return c
}

func (c *LoggerOptions) UseDefaultNameKeyIfNotSpecified() *LoggerOptions {
	if len(c.NameKey) == 0 {
		c.NameKey = defaultNameKey
	}
	return c
}

func (c *LoggerOptions) UseDefaultCallerKeyIfNotSpecified() *LoggerOptions {
	if len(c.CallerKey) == 0 {
		c.CallerKey = defaultCallerKey
	}
	return c
}

func (c *LoggerOptions) UseDefaultMessageKeyIfNotSpecified() *LoggerOptions {
	if len(c.MessageKey) == 0 {
		c.MessageKey = defaultMessageKey
	}
	return c
}
func (c *LoggerOptions) UseDefaultFunctionKeyIfNotSpecified() *LoggerOptions {
	if len(c.FunctionKey) == 0 {
		c.FunctionKey = defaultFunctionKey
	}
	return c
}

func (c *LoggerOptions) UseDefaultStackTraceKeyIfNotSpecified() *LoggerOptions {
	if len(c.StackTraceKey) == 0 {
		c.StackTraceKey = defaultStackTraceKey
	}
	return c
}

func (c *LoggerOptions) UseDefaultEncodeLevelIfNotSpecified() *LoggerOptions {
	if c.EncodeLevel == nil {
		c.EncodeLevel = &defaultEncodeLevel
	}
	return c
}

func (c *LoggerOptions) UseDefaultEncodeTimeIfNotSpecified() *LoggerOptions {
	if c.EncodeTime == nil {
		c.EncodeTime = &defaultEncodeTime
	}
	return c
}

func (c *LoggerOptions) UseDefaultEncodeDurationIfNotSpecified() *LoggerOptions {
	if c.EncodeDuration == nil {
		c.EncodeDuration = &defaultEncodeDuration
	}
	return c
}

func (c *LoggerOptions) UseDefaultEncodeCallerIfNotSpecified() *LoggerOptions {
	if c.EncodeCaller == nil {
		c.EncodeCaller = &defaultEncodeCaller
	}
	return c
}
