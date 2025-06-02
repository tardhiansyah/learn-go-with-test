package find

type Person struct {
	Name string
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return value, false
}
