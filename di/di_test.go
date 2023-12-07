package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	actual := buffer.String()
	expected := "Hello, Chris"

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}
