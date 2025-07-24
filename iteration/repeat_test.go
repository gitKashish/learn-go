package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat character more than 1 time", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"
		assertSameResult(t, repeated, expected)
	})

	t.Run("repeat character 0 times", func(t *testing.T) {
		repeated := Repeat("b", 0)
		expected := ""
		assertSameResult(t, repeated, expected)
	})
}

func assertSameResult(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 8)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 4)
	fmt.Println(repeated)
	// Output: cccc
}
