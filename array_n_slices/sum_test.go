package arraynslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("find sum of collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertSameResult(t, got, want, numbers)

	})
}

func TestSumAll(t *testing.T) {
	t.Run("find sum of multiple arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("find sum of all numbers in array except head", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 3, 9})
		want := []int{2, 12}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty arrays", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func assertSameResult(t testing.TB, got, want int, value []int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d but go %d, given %v", want, got, value)
	}
}

func checkSums(t testing.TB, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
