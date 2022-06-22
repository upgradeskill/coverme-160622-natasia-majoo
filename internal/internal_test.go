package internal

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("DoesSomething/It actually does something and returns 5", func(t *testing.T) {
		got := DoesSomethingAndReturn5()
		want := uint(5)

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
}
