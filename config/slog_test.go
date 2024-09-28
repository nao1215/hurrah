package config

import (
	"bytes"
	"testing"
)

func TestNewStructuredLogger(t *testing.T) {
	t.Run("debug level message", func(t *testing.T) {
		w := new(bytes.Buffer)
		logger := NewStructuredLogger(w, true)
		if logger == nil {
			t.Error("logger is nil")
		}

		logger.Debug("debug-message")
		for _, want := range []string{
			"msg\":\"debug-message\"",
			"level\":\"DEBUG\"",
		} {
			if !bytes.Contains(w.Bytes(), []byte(want)) {
				t.Errorf("want %q to be contained in %q", want, w.String())
			}
		}
	})

	t.Run("info level message", func(t *testing.T) {
		w := new(bytes.Buffer)
		logger := NewStructuredLogger(w, true)
		if logger == nil {
			t.Error("logger is nil")
		}

		logger.Info("info-message")
		for _, want := range []string{
			"msg\":\"info-message\"",
			"level\":\"INFO\"",
		} {
			if !bytes.Contains(w.Bytes(), []byte(want)) {
				t.Errorf("want %q to be contained in %q", want, w.String())
			}
		}
	})

	t.Run("warn level message", func(t *testing.T) {
		w := new(bytes.Buffer)
		logger := NewStructuredLogger(w, true)
		if logger == nil {
			t.Error("logger is nil")
		}

		logger.Warn("warn-message")
		for _, want := range []string{
			"msg\":\"warn-message\"",
			"level\":\"WARN\"",
		} {
			if !bytes.Contains(w.Bytes(), []byte(want)) {
				t.Errorf("want %q to be contained in %q", want, w.String())
			}
		}
	})

	t.Run("error level message", func(t *testing.T) {
		w := new(bytes.Buffer)
		logger := NewStructuredLogger(w, true)
		if logger == nil {
			t.Error("logger is nil")
		}

		logger.Error("error-message")
		for _, want := range []string{
			"msg\":\"error-message\"",
			"level\":\"ERROR\"",
		} {
			if !bytes.Contains(w.Bytes(), []byte(want)) {
				t.Errorf("want %q to be contained in %q", want, w.String())
			}
		}
	})
}
