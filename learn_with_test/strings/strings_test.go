package strings

import (
	"testing"
)

func TestClone(t *testing.T) {
	t.Run("clone string", func(t *testing.T) {
		original := "hello"
		cloned := Clone(original)
		expected := "hello"

		if cloned != expected {
			t.Errorf("got %q, want %q", cloned, expected)
		}

		if &cloned != &original {
			t.Errorf("cloned addr %p, original addr %p", &cloned, &original)
		}
	})
}

