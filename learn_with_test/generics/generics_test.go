package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})

	t.Run("asserting integer stack", func(t *testing.T) {
		stack := NewStack[int]()

		AssertTrue(t, stack.IsEmpty())

		stack.Push(1)
		AssertFalse(t, stack.IsEmpty())

		stack.Push(2)
		value, _ := stack.Pop()
		AssertEqual(t, value, 2)
		value, _ = stack.Pop()
		AssertEqual(t, value, 1)
		AssertTrue(t, stack.IsEmpty())
	})

	t.Run("asserting string stack", func(t *testing.T) {
		stack := NewStack[string]()

		AssertTrue(t, stack.IsEmpty())

		stack.Push("hello")
		AssertFalse(t, stack.IsEmpty())

		stack.Push("world")
		value, _ := stack.Pop()
		AssertEqual(t, value, "world")
		value, _ = stack.Pop()
		AssertEqual(t, value, "hello")
		AssertTrue(t, stack.IsEmpty())
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("got %v, want not %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
