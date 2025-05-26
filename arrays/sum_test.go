package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of all slices", func(t *testing.T) {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3, 9}

		// DeepEqual similar to creating a function that iterates over the slices
		// reflect.DeepEqual isn't type safe, we can use slices.Equal but it's only for comparable types
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("Sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("Sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 6})
		want := []int{0, 10}

		checkSums(t, got, want)
	})

}