package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Charlie")
	got := buffer.String()
	want := "hello Charlie"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
