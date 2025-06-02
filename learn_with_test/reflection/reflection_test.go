package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	// Table-based tests
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "Gotham"},
			[]string{"Chris", "Gotham"},
		},
		{
			"strunct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 32},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{
				Name: "Chris",
				Profile: Profile{
					Age:  32,
					City: "Gotham",
				},
			},
			[]string{"Chris", "Gotham"},
		},
		{
			"pointer to things",
			&Person{
				Name: "Chris",
				Profile: Profile{
					Age:  32,
					City: "Gotham",
				},
			},
			[]string{"Chris", "Gotham"},
		},
		{
			"slice of things",
			[]Profile{
				{Age: 32, City: "Gotham"},
				{Age: 28, City: "Metropolis"},
			},
			[]string{"Gotham", "Metropolis"},
		},
		{
			"array of things",
			[2]Profile{
				{Age: 32, City: "Gotham"},
				{Age: 28, City: "Metropolis"},
			},
			[]string{"Gotham", "Metropolis"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Chris": "Gotham",
			"Bruce": "Metropolis",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Gotham")
		assertContains(t, got, "Metropolis")
	})

	t.Run("nested maps", func(t *testing.T) {
		aMap := map[string]map[string]string{
			"Chris": {
				"City": "Gotham",
				"Age":  "32",
			},
			"Bruce": {
				"City": "Metropolis",
				"Age":  "28",
			},
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Gotham")
		assertContains(t, got, "Metropolis")
		assertContains(t, got, "32")
		assertContains(t, got, "28")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Gotham"}
			aChannel <- Profile{28, "Metropolis"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Gotham", "Metropolis"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{Age: 30, City: "Gotham"}, Profile{Age: 25, City: "Metropolis"}
		}

		var got []string
		want := []string{"Gotham", "Metropolis"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, item := range haystack {
		if item == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected to find %q but didn't in %v", needle, haystack)
	}
}
