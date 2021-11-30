package options

import (
	"jarjarbinks/pkg/infrastructure/logging/utilities"
	"testing"
)

func TestNewDefaultLoggerOptionsShouldReturnDefaultLoggerOptions(t *testing.T) {
	defaultOptions := New()
	if defaultOptions == nil {
		t.Error("default option could not be nil")
	}
}

func TestNewDefaultLoggerOptionsShouldDevelopmentModeReturnFalseWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.Development {
		t.Error("development mode could not be selected as a default option")
	}
}

func TestNewDefaultLoggerOptionsShouldEncodingMustEqualsToDefaultEncodingWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.Encoding) == 0 {
		t.Error("Encoding could not be nil or empty")
	}

	if defaultOptions.Encoding != defaultEncoding {
		t.Errorf("%s must be a default encoding", defaultEncoding)
	}
}
func TestNewDefaultLoggerOptionsShouldOutputPathsMustEqualsToDefaultOutputPathsWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.OutputPaths == nil || len(defaultOptions.OutputPaths) == 0 {
		t.Error("OutputPaths could not be nil or empty")
	}

	if !utilities.CompareSlices(defaultOptions.OutputPaths, defaultOutputPaths) {
		t.Error("OutputPaths must be equals default OutputPaths")
	}
}

func TestNewDefaultLoggerOptionsShouldErrorOutputPathsMustEqualsToDefaultErrorOutputPathsWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.ErrorOutputPaths == nil || len(defaultOptions.ErrorOutputPaths) == 0 {
		t.Error("ErrorOutputPaths could not be nil or empty")
	}

	if !utilities.CompareSlices(defaultOptions.ErrorOutputPaths, defaultErrorOutputPaths) {
		t.Error("ErrorOutputPaths must be equals default ErrorOutputPaths")
	}
}

func TestNewDefaultLoggerOptionsShouldLoggingLevelMustEqualsToDefaultLoggingLevelWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.Level) == 0 {
		t.Error("Level could not be nil or empty")
	}

	if defaultOptions.Level != defaultLoggingLevel {
		t.Errorf("%s must be specified as a default Level", defaultLoggingLevel)
	}
}

func TestNewDefaultLoggerOptionsShouldTimeKeyMustEqualsToDefaultTimeKeyWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.TimeKey) == 0 {
		t.Error("TimeKey could not be nil or empty")
	}

	if defaultOptions.TimeKey != defaultTimeKey {
		t.Errorf("%s must be specified as a default TimeKey", defaultTimeKey)
	}
}

func TestNewDefaultLoggerOptionsShouldLevelKeyMustEqualsToDefaultLevelKeyWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.LevelKey) == 0 {
		t.Error("LevelKey  could not be nil or empty")
	}

	if defaultOptions.LevelKey != defaultLevelKey {
		t.Errorf("%s must be specified as a default LevelKey", defaultLevelKey)
	}
}

func TestNewDefaultLoggerOptionsShouldNameKeyMustEqualsToDefaultNameKeyWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.NameKey) == 0 {
		t.Error("NameKey  could not be nil or empty")
	}

	if defaultOptions.NameKey != defaultNameKey {
		t.Errorf("%s must be specified as a default NameKey", defaultNameKey)
	}
}

func TestNewDefaultLoggerOptionsShouldCallerKeyMustEqualsToDefaultCallerKeyWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.CallerKey) == 0 {
		t.Error("CallerKey could not be nil or empty")
	}

	if defaultOptions.CallerKey != defaultCallerKey {
		t.Errorf("%s must be specified as a default CallerKey", defaultCallerKey)
	}
}

func TestNewDefaultLoggerOptionsShouldMessageKeyMustEqualsToDefaultMessageKeyWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.MessageKey) == 0 {
		t.Error("MessageKey could not be nil or empty")
	}

	if defaultOptions.MessageKey != defaultMessageKey {
		t.Errorf("%s must be specified as a default MessageKey", defaultMessageKey)
	}
}

func TestNewDefaultLoggerOptionsShouldStackTraceKeyMustEqualsToDefaultStackTraceWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if len(defaultOptions.StackTraceKey) == 0 {
		t.Error("StackTraceKey could not be nil or empty")
	}

	if defaultOptions.StackTraceKey != defaultStackTraceKey {
		t.Errorf("%s must be specified as a default StackTraceKey", defaultStackTraceKey)
	}
}

func TestNewDefaultLoggerOptionsShouldEncodeLevelMustEqualsToDefaultEncodeLevelWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.EncodeLevel == nil {
		t.Error("EncodeLevel could not be nil or empty")
	}

	if defaultOptions.EncodeLevel != &defaultEncodeLevel {
		t.Errorf("%d must be specified as a default EncodeLevel", &defaultEncodeLevel)
	}
}

func TestNewDefaultLoggerOptionsShouldEncodeTimeMustEqualsToDefaultEncodeTimeWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.EncodeTime == nil {
		t.Error("EncodeTime could not be nil or empty")
	}

	if defaultOptions.EncodeTime != &defaultEncodeTime {
		t.Errorf("%d must be specified as a default EncodeTime", &defaultEncodeTime)
	}
}

func TestNewDefaultLoggerOptionsShouldEncodeCallerMustEqualsToDefaultEncodeCallerWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.EncodeCaller == nil {
		t.Error("EncodeCaller could not be nil or empty")
	}

	if defaultOptions.EncodeCaller != &defaultEncodeCaller {
		t.Errorf("%d must be specified as a default EncodeCaller", &defaultEncodeCaller)
	}
}

func TestNewDefaultLoggerOptionsShouldEncodeDurationMustEqualsToDefaultEncodeDurationWhenDefaultOptionsSelected(t *testing.T) {
	defaultOptions := New()
	if defaultOptions.EncodeDuration == nil {
		t.Error("EncodeDuration could not be nil or empty")
	}

	if defaultOptions.EncodeDuration != &defaultEncodeDuration {
		t.Errorf("%d must be specified as a default EncodeDuration", &defaultEncodeDuration)
	}
}

func TestUseDefaultEncodingIfNotSpecifiedShouldReturnDefaultEncodingWhenEncodingDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultEncodingIfNotSpecified()

	if len(options.Encoding) == 0 {
		t.Error("Encoding could not be nil or empty")
	}

	if options.Encoding != defaultEncoding {
		t.Errorf("%s must be specified as a default Encoding", defaultEncoding)
	}

}

func TestUseDefaultOutputPathsIfNotSpecifiedShouldReturnDefaultOutputPathsWhenOutputPathsDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultOutputPathsIfNotSpecified()

	if options.OutputPaths == nil || len(options.OutputPaths) == 0 {
		t.Error("OutputPaths could not be nil or empty")
	}

	if !utilities.CompareSlices(options.OutputPaths, defaultOutputPaths) {
		t.Error("OutputPaths must be equals default OutputPaths")
	}

}

func TestUseDefaultErrorOutputPathsIfNotSpecifiedShouldReturnDefaultErrorOutputPathsWhenErrorOutputPathsDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultErrorOutputPathsIfNotSpecified()

	if options.ErrorOutputPaths == nil || len(options.ErrorOutputPaths) == 0 {
		t.Error("ErrorOutputPaths could not be nil or empty")
	}

	if !utilities.CompareSlices(options.ErrorOutputPaths, defaultErrorOutputPaths) {
		t.Error("ErrorOutputPaths must be equals default ErrorOutputPaths")
	}

}

func TestUseDefaultLoggingLevelIfNotSpecifiedShouldReturnDefaultLoggingLevelWhenLoggingLevelDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultLoggingLevelIfNotSpecified()

	if len(options.Level) == 0 {
		t.Error("Level could not be nil or empty")
	}

	if options.Level != defaultLoggingLevel {
		t.Errorf("%s must be specified as a default Level", defaultLoggingLevel)
	}

}

func TestUseDefaultTimeKeyIfNotSpecifiedShouldReturnDefaultTimeKeyWhenTimeKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultTimeKeyIfNotSpecified()

	if len(options.TimeKey) == 0 {
		t.Error("TimeKey could not be nil or empty")
	}

	if options.TimeKey != defaultTimeKey {
		t.Errorf("%s must be specified as a default TimeKey", defaultTimeKey)
	}
}

func TestUseDefaultLevelKeyIfNotSpecifiedShouldReturnDefaultLevelKeyWhenTimeKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultLevelKeyIfNotSpecified()

	if len(options.LevelKey) == 0 {
		t.Error("LevelKey could not be nil or empty")
	}

	if options.LevelKey != defaultLevelKey {
		t.Errorf("%s must be specified as a default LevelKey", defaultLevelKey)
	}
}

func TestUseDefaultNameKeyIfNotSpecifiedShouldReturnDefaultNameKeyWhenNameKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultNameKeyIfNotSpecified()

	if len(options.NameKey) == 0 {
		t.Error("TimeKey could not be nil or empty")
	}

	if options.NameKey != defaultNameKey {
		t.Errorf("%s must be specified as a default NameKey", defaultNameKey)
	}
}

func TestUseDefaultCallerKeyIfNotSpecifiedShouldReturnDefaultNameKeyWhenCallerKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultCallerKeyIfNotSpecified()

	if len(options.CallerKey) == 0 {
		t.Error("CallerKey could not be nil or empty")
	}

	if options.CallerKey != defaultCallerKey {
		t.Errorf("%s must be specified as a default CallerKey", defaultCallerKey)
	}
}

func TestUseDefaultMessageKeyIfNotSpecifiedShouldReturnDefaultNameKeyWhenMessageKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultMessageKeyIfNotSpecified()

	if len(options.MessageKey) == 0 {
		t.Error("MessageKey could not be nil or empty")
	}

	if options.MessageKey != defaultMessageKey {
		t.Errorf("%s must be specified as a default MessageKey", defaultMessageKey)
	}
}

func TestUseDefaultStackTraceKeyIfNotSpecifiedShouldReturnDefaultNameKeyWhenStackTraceKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultStackTraceKeyIfNotSpecified()

	if len(options.StackTraceKey) == 0 {
		t.Error("StackTraceKey could not be nil or empty")
	}

	if options.StackTraceKey != defaultStackTraceKey {
		t.Errorf("%s must be specified as a default StackTraceKey", defaultStackTraceKey)
	}
}

func TestUseDefaultEncodeLevelIfNotSpecifiedShouldReturnDefaultEncodeLevelWhenStackTraceKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultEncodeLevelIfNotSpecified()

	if options.EncodeLevel == nil {
		t.Error("EncodeLevel could not be nil or empty")
	}

	if options.EncodeLevel != &defaultEncodeLevel {
		t.Errorf("%d must be specified as a default EncodeLevel", &defaultEncodeLevel)
	}
}

func TestUseDefaultEncodeTimeIfNotSpecifiedShouldReturnDefaultEncodeTimeWhenStackTraceKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultEncodeTimeIfNotSpecified()

	if options.EncodeTime == nil {
		t.Error("EncodeTime could not be nil or empty")
	}

	if options.EncodeTime != &defaultEncodeTime {
		t.Errorf("%d must be specified as a default EncodeTime", &defaultEncodeTime)
	}
}

func TestUseDefaultEncodeDurationIfNotSpecifiedShouldReturnDefaultEncodeDurationWhenStackTraceKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultEncodeDurationIfNotSpecified()

	if options.EncodeDuration == nil {
		t.Error("EncodeDuration could not be nil or empty")
	}

	if options.EncodeDuration != &defaultEncodeDuration {
		t.Errorf("%d must be specified as a default EncodeDuration", &defaultEncodeDuration)
	}
}

func TestUseDefaultEncodeCallerIfNotSpecifiedShouldReturnDefaultEncodeCallerWhenStackTraceKeyDoesNotSelected(t *testing.T) {
	options := &LoggerOptions{}
	options.UseDefaultEncodeCallerIfNotSpecified()

	if options.EncodeCaller == nil {
		t.Error("EncodeCaller could not be nil or empty")
	}

	if options.EncodeCaller != &defaultEncodeCaller {
		t.Errorf("%d must be specified as a default EncodeCaller", &defaultEncodeCaller)
	}
}
