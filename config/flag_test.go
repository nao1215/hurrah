package config

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFlag(t *testing.T) {
	t.Run("NewFlag without argument. Flag has default value", func(t *testing.T) {
		args := []string{"cmd/hurrah/main_test.go"}

		orgArgs := os.Args
		defer func() { os.Args = orgArgs }()
		os.Args = args

		got := NewFlag()
		want := Flag{Port: 8080, ConfigFile: "config.yaml", Debug: false}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("NewFlag() mismatch (-got +want):\n%s", diff)
		}
	})

	t.Run("NewFlag with arguments. Flag has the given value", func(t *testing.T) {
		args := []string{"cmd/hurrah/main_test.go", "-port", "1234", "-config", "new.yaml", "-debug"}

		orgArgs := os.Args
		defer func() { os.Args = orgArgs }()
		os.Args = args

		got := NewFlag()
		want := Flag{Port: 1234, ConfigFile: "new.yaml", Debug: true}

		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("NewFlag() mismatch (-got +want):\n%s", diff)
		}
	})
}
