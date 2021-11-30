package utilities

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jarjarbinks/pkg/infrastructure/logging/enumeration"
	"reflect"
	"testing"
)

func TestToZapFieldsShouldReturnEmptyZapFieldsWhenParametersIsNil(t *testing.T) {
	zapFields := ToZapFields()
	if zapFields == nil {
		t.Error("ZapFields option could not be nil")
	}
}

func TestToZapFieldsShouldReturnEmptyZapFieldsWhenParametersIsEmpty(t *testing.T) {
	var parameters map[string]interface{}
	zapFields := ToZapFields(parameters)
	if zapFields == nil {
		t.Error("ZapFields option could not be nil")
	}
}

func TestToZapFieldsShouldReturnZapFieldsWhenParametersDoesNotEmpty(t *testing.T) {
	parameters := map[string]interface{}{
		"sample":   "1",
		"sample_2": 1,
		"sample_3": 1.1,
		"sample_4": true,
		"":         true,
	}
	zapFields := ToZapFields(parameters)
	if zapFields == nil {
		t.Error("ZapFields option could not be nil")
	}

	for _, zapField := range *zapFields {
		switch zapField.Key {
		case "sample":
			if zapField.Type != zapcore.StringType {
				t.Errorf("%s field could not be recognized. it must be StringType", zapField.Key)
			}
		case "sample_2":
			if zapField.Type != zapcore.Int64Type {
				t.Errorf("%s field could not be recognized. it must be Int64Type", zapField.Key)
			}
		case "sample_3":
			if zapField.Type != zapcore.Float64Type {
				t.Errorf("%s field could not be recognized. it must be Float64Type", zapField.Key)
			}
		case "sample_4":
			if zapField.Type != zapcore.BoolType {
				t.Errorf("%s field could not be recognized. it must be BoolType", zapField.Key)
			}
		}
	}
}

func TestToZapLogLevelShouldReturnDEBUGWhenSelectedLevelEqualsToDEBUG(t *testing.T) {
	zapLevel := ToZapLogLevel("DEBUG")
	if zapLevel != zapcore.DebugLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapLogLevelShouldReturnINFOWhenSelectedLevelEqualsToINFO(t *testing.T) {
	zapLevel := ToZapLogLevel("INFO")
	if zapLevel != zapcore.InfoLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapLogLevelShouldReturnWARNWhenSelectedLevelEqualsToWARN(t *testing.T) {
	zapLevel := ToZapLogLevel("WARN")
	if zapLevel != zapcore.WarnLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapLogLevelShouldReturnERRORWhenSelectedLevelEqualsToERROR(t *testing.T) {
	zapLevel := ToZapLogLevel("ERROR")
	if zapLevel != zapcore.ErrorLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapLogLevelShouldReturnPANICWhenSelectedLevelEqualsToPANIC(t *testing.T) {
	zapLevel := ToZapLogLevel("PANIC")
	if zapLevel != zapcore.PanicLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapLogLevelShouldReturnFATALWhenSelectedLevelEqualsToFATAL(t *testing.T) {
	zapLevel := ToZapLogLevel("FATAL")
	if zapLevel != zapcore.FatalLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapLogLevelShouldReturnDEBUGWhenSelectedLevelDoesNotEqualsDefinedLevels(t *testing.T) {
	zapLevel := ToZapLogLevel("NOT_DEFINED_LEVEL")
	if zapLevel != zapcore.InfoLevel {
		t.Error("zapLevel could not be recognized")
	}
}

func TestToZapFieldShouldReturnIntWhenValueEqualsToInteger(t *testing.T) {
	zapField, err := toZapField("sample", 1)
	if err != nil {
		t.Error("zapField must be valid")
	}

	if zapField != zap.Int("sample", 1) {
		t.Error("zapField must be Int")
	}
}

func TestToZapFieldShouldReturnFloatWhenValueEqualsToFloat(t *testing.T) {
	zapField, err := toZapField("sample", 1.2)
	if err != nil {
		t.Error("zapField must be valid")
	}
	if zapField != zap.Float64("sample", 1.2) {
		t.Error("zapField must be Float64")
	}
}

func TestToZapFieldShouldReturnStringWhenValueEqualsToString(t *testing.T) {
	zapField, err := toZapField("sample", "sample")
	if err != nil {
		t.Error("zapField must be valid")
	}
	if zapField != zap.String("sample", "sample") {
		t.Error("zapField must be ExtractBodyAsString")
	}
}

func TestToZapFieldShouldReturnBooleanWhenValueEqualsToBoolean(t *testing.T) {
	zapField, err := toZapField("sample", true)
	if err != nil {
		t.Error("zapField must be valid")
	}
	if zapField != zap.Bool("sample", true) {
		t.Error("zapField must be Boolean")
	}
}

func TestToZapFieldShouldReturnErrorWhenKeyIsEmpty(t *testing.T) {
	_, err := toZapField("", true)
	if err == nil {
		t.Error("zapField must have an error when key is nil")
	}
}

func TestToZapLevelEncoderShouldReturnLowercaseWhenLowercaseSelected(t *testing.T) {
	levelEncoder := ToZapLevelEncoder(enumeration.Lowercase)
	if reflect.TypeOf(levelEncoder).Size() != reflect.TypeOf(zapcore.LowercaseLevelEncoder).Size() {
		t.Error("level encoder must be equals to LowercaseLevelEncoder")
	}
}

func TestToZapLevelEncoderShouldReturnCamelcaseWhenCamelcaseSelected(t *testing.T) {
	levelEncoder := ToZapLevelEncoder(enumeration.Camelcase)
	if reflect.TypeOf(levelEncoder).Size() != reflect.TypeOf(zapcore.CapitalLevelEncoder).Size() {
		t.Error("level encoder must be equals to CapitalLevelEncoder")
	}
}

func TestToZapLevelEncoderShouldReturnLowercaseWhenDefaultSelected(t *testing.T) {
	levelEncoder := ToZapLevelEncoder(enumeration.DefaultLevelEncoder)
	if reflect.TypeOf(levelEncoder).Size() != reflect.TypeOf(zapcore.LowercaseLevelEncoder).Size() {
		t.Error("level encoder must be equals to LowercaseLevelEncoder")
	}
}

func TestToZapTimeEncoderShouldReturnRFC3339NanoWhenRFC3339NanoSelected(t *testing.T) {
	timeEncoder := ToZapTimeEncoder(enumeration.RFC3339Nano)
	if reflect.TypeOf(timeEncoder).Size() != reflect.TypeOf(zapcore.RFC3339NanoTimeEncoder).Size() {
		t.Error("time encoder must be equals to RFC3339NanoTimeEncoder")
	}
}

func TestToZapTimeEncoderShouldReturnRFC3339WhenRFC3339Selected(t *testing.T) {
	timeEncoder := ToZapTimeEncoder(enumeration.RFC3339)
	if reflect.TypeOf(timeEncoder).Size() != reflect.TypeOf(zapcore.RFC3339TimeEncoder).Size() {
		t.Error("time encoder must be equals to RFC3339TimeEncoder")
	}
}

func TestToZapTimeEncoderShouldReturnISO8601WhenISO8601Selected(t *testing.T) {
	timeEncoder := ToZapTimeEncoder(enumeration.ISO8601)
	if reflect.TypeOf(timeEncoder).Size() != reflect.TypeOf(zapcore.ISO8601TimeEncoder).Size() {
		t.Error("time encoder must be equals to ISO8601TimeEncoder")
	}
}

func TestToZapTimeEncoderShouldReturnMillisecondsWhenMillisecondsSelected(t *testing.T) {
	timeEncoder := ToZapTimeEncoder(enumeration.Milliseconds)
	if reflect.TypeOf(timeEncoder).Size() != reflect.TypeOf(zapcore.EpochMillisTimeEncoder).Size() {
		t.Error("time encoder must be equals to EpochMillisTimeEncoder")
	}
}

func TestToZapTimeEncoderShouldReturnNanosecondsWhenNanosecondsSelected(t *testing.T) {
	timeEncoder := ToZapTimeEncoder(enumeration.Nanoseconds)
	if reflect.TypeOf(timeEncoder).Size() != reflect.TypeOf(zapcore.EpochNanosTimeEncoder).Size() {
		t.Error("time encoder must be equals to EpochNanosTimeEncoder")
	}
}

func TestToZapTimeEncoderShouldReturnEpochTimeWhenDefaultSelected(t *testing.T) {
	timeEncoder := ToZapTimeEncoder(enumeration.DefaultTimeEncoder)
	if reflect.TypeOf(timeEncoder).Size() != reflect.TypeOf(zapcore.EpochTimeEncoder).Size() {
		t.Error("time encoder must be equals to EpochTimeEncoder")
	}
}

func TestToZapDurationEncoderShouldReturnStringDurationWhenStringDurationSelected(t *testing.T) {
	durationEncoder := ToZapDurationEncoder(enumeration.StringDuration)
	if reflect.TypeOf(durationEncoder).Size() != reflect.TypeOf(zapcore.StringDurationEncoder).Size() {
		t.Error("duration encoder must be equals to StringDurationEncoder")
	}
}

func TestToZapDurationEncoderShouldReturnMillisecondDurationWhenMillisecondDurationSelected(t *testing.T) {
	durationEncoder := ToZapDurationEncoder(enumeration.MillisecondDuration)
	if reflect.TypeOf(durationEncoder).Size() != reflect.TypeOf(zapcore.MillisDurationEncoder).Size() {
		t.Error("time encoder must be equals to MillisDurationEncoder")
	}
}

func TestToZapDurationEncoderShouldReturnNanosecondDurationWhenNanosecondDurationSelected(t *testing.T) {
	durationEncoder := ToZapDurationEncoder(enumeration.NanosecondDuration)
	if reflect.TypeOf(durationEncoder).Size() != reflect.TypeOf(zapcore.NanosDurationEncoder).Size() {
		t.Error("duration encoder must be equals to NanosDurationEncoder")
	}
}

func TestToZapDurationEncoderShouldReturnSecondsDurationWhenDefaultDurationSelected(t *testing.T) {
	durationEncoder := ToZapDurationEncoder(enumeration.DefaultDurationEncoder)
	if reflect.TypeOf(durationEncoder).Size() != reflect.TypeOf(zapcore.SecondsDurationEncoder).Size() {
		t.Error("duration encoder must be equals to SecondsDurationEncoder")
	}
}

func TestToZapCallerEncoderShouldReturnLongestFunctionNameWhenLongestFunctionNameSelected(t *testing.T) {
	callerEncoder := ToZapCallerEncoder(enumeration.LongestFunctionName)
	if reflect.TypeOf(callerEncoder).Size() != reflect.TypeOf(zapcore.FullCallerEncoder).Size() {
		t.Error("caller encoder must be equals to FullCallerEncoder")
	}
}

func TestToZapCallerEncoderShouldReturnShortestFunctionNameWhenShortestFunctionNameSelected(t *testing.T) {
	callerEncoder := ToZapCallerEncoder(enumeration.ShortestFunctionName)
	if reflect.TypeOf(callerEncoder).Size() != reflect.TypeOf(zapcore.ShortCallerEncoder).Size() {
		t.Error("caller encoder must be equals to ShortCallerEncoder")
	}
}

func TestToZapCallerEncoderShouldReturnLongestFunctionNameWhenDefaultSelected(t *testing.T) {
	callerEncoder := ToZapCallerEncoder(enumeration.DefaultCallerEncoder)
	if reflect.TypeOf(callerEncoder).Size() != reflect.TypeOf(zapcore.FullCallerEncoder).Size() {
		t.Error("caller encoder must be equals to FullCallerEncoder")
	}
}
