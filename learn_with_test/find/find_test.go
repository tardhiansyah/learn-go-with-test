package find

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 3, 5, 4, 2}
		expected := 4

		firstEvenNumber, found := Find(numbers, func(n int) bool {
			return n%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, expected)
	})

	t.Run("find first person with first name 'Alice'", func(t *testing.T) {
		people := []Person{
			{Name: "Bob Smith"},
			{Name: "Alice Johnson"},
			{Name: "Charlie Alice"},
		}
		expected := Person{Name: "Alice Johnson"}

		firstPerson, found := Find(people, func(p Person) bool {
			return strings.HasPrefix(p.Name, "Alice")
		})

		AssertTrue(t, found)
		AssertEqual(t, firstPerson, expected)
	})
}

func AssertTrue(t *testing.T, condition bool) {
	t.Helper()
	if !condition {
		t.Error("expected condition to be true")
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
