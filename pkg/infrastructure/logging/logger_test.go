package logging

import (
	"errors"
	"io/ioutil"
	"jarjarbinks/pkg/infrastructure/logging/options"
	"os"
	"testing"
)

func TestNewLoggerShouldReturnLoggerWithDefaultOptionsWhenOptionsDoesNotSelected(t *testing.T) {
	defaultLogger, err := New()
	if err != nil {
		t.Error("error does not expected when default-logger initialized")
	}

	if defaultLogger == nil {
		t.Errorf("default-logger must be initialized")
	}

}

func TestNewLoggerWithOptionsShouldReturnLoggerWithSelectedOptionsWhenOptionsSelected(t *testing.T) {
	defaultLogger, err := New(options.New())
	if err != nil {
		t.Error("error does not expected when default-logger initialized")
	}

	if defaultLogger == nil {
		t.Errorf("default-logger must be initialized")
	}
}

func TestInfoShouldBuildInfoLog(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Info("sample")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("info log could not be written")
	}
}

func TestInfoShouldBuildInfoLogWithParameters(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Info("sample", map[string]interface{}{"sample": "foo"})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("info log could not be written")
	}
}

func TestDebugShouldBuildDebugLog(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Debug("sample")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("Debug log could not be written")
	}
}

func TestDebugShouldBuildDebugLogWithParameters(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Debug("sample", map[string]interface{}{"sample": "foo"})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("Debug log could not be written")
	}
}

func TestWarnShouldBuildWarnLog(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Warn("sample")
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("Warn log could not be written")
	}
}

func TestWarnShouldBuildWarnLogWithParameters(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Warn("sample", map[string]interface{}{"sample": "foo"})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("Warn log could not be written")
	}
}

func TestErrorShouldBuildErrorLog(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Error("sample", errors.New("sample"))
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("ExtractError log could not be written")
	}
}

func TestErrorWithParametersShouldBuildErrorLog(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Error("sample", errors.New("sample"), map[string]interface{}{"sample": "foo"})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("ExtractError log could not be written")
	}
}

func TestErrorWithParametersShouldBuildErrorLogWhenErrorIsNil(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defaultLogger, _ := New(options.New())
	defaultLogger.Error("sample", nil)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if len(out) == 0 {
		t.Error("ExtractError log could not be written")
	}
}
