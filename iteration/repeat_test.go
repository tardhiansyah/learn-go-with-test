package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat single character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat character 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat multiple characters 5 times", func(t *testing.T) {
		repeated := Repeat("abc", 5)
		expected := "abcabcabcabcabc"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat multiple characters negative times", func(t *testing.T) {
		repeated := Repeat("abc", -5)
		expected := ""

		assertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}