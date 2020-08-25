package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectSum := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	}

	t.Run("Sum of an array of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want)
	})

	t.Run("Sum of a collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertCorrectSum(t, got, want)
	})

}

func TestSumAll(t *testing.T) {
	assertCorrectSumAll := func(t *testing.T, got, want []int) {
		t.Helper()

		// reflect.DeepEqual is NOT TYPE SAFE.

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v wanted %v", got, want)
		}
	}

	t.Run("Sum of a collection of collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorrectSumAll(t, got, want)
	})

	t.Run("Sum of tails of all the collections", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertCorrectSumAll(t, got, want)
	})

	t.Run("Safely sum empty tail slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrectSumAll(t, got, want)
	})
}
