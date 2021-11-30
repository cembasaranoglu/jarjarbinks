package providers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jarjarbinks/pkg/infrastructure/logging/options"
	"jarjarbinks/pkg/infrastructure/logging/utilities"
)

func New(configuration ...*options.LoggerOptions) (*zap.Logger, error) {
	logger, err := newZapConfig(configuration...).Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return logger, nil
}

func newZapConfig(loggerOptions ...*options.LoggerOptions) *zap.Config {
	var loggerOption options.LoggerOptions
	if loggerOptions == nil {
		loggerOption = *options.New()
	} else {
		loggerOption = *loggerOptions[0]
	}
	loggerOption.
		UseDefaultEncodingIfNotSpecified().
		UseDefaultOutputPathsIfNotSpecified().
		UseDefaultErrorOutputPathsIfNotSpecified().
		UseDefaultLoggingLevelIfNotSpecified().
		UseDefaultTimeKeyIfNotSpecified().
		UseDefaultLevelKeyIfNotSpecified().
		UseDefaultNameKeyIfNotSpecified().
		UseDefaultCallerKeyIfNotSpecified().
		UseDefaultMessageKeyIfNotSpecified().
		UseDefaultStackTraceKeyIfNotSpecified().
		UseDefaultEncodeLevelIfNotSpecified().
		UseDefaultEncodeTimeIfNotSpecified().
		UseDefaultEncodeDurationIfNotSpecified().
		UseDefaultEncodeCallerIfNotSpecified().UseDefaultFunctionKeyIfNotSpecified()

	return &zap.Config{
		Level:       zap.NewAtomicLevelAt(utilities.ToZapLogLevel(loggerOption.Level)),
		Development: loggerOption.Development,
		Encoding:    loggerOption.Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        loggerOption.TimeKey,
			LevelKey:       loggerOption.LevelKey,
			NameKey:        loggerOption.NameKey,
			CallerKey:      loggerOption.CallerKey,
			FunctionKey:    loggerOption.FunctionKey,
			MessageKey:     loggerOption.MessageKey,
			StacktraceKey:  loggerOption.StackTraceKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    utilities.ToZapLevelEncoder(*loggerOption.EncodeLevel),
			EncodeTime:     utilities.ToZapTimeEncoder(*loggerOption.EncodeTime),
			EncodeDuration: utilities.ToZapDurationEncoder(*loggerOption.EncodeDuration),
			EncodeCaller:   utilities.ToZapCallerEncoder(*loggerOption.EncodeCaller),
		},
		InitialFields:    loggerOption.DefaultParameters,
		OutputPaths:      loggerOption.OutputPaths,
		ErrorOutputPaths: loggerOption.ErrorOutputPaths,
	}
}
