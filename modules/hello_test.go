package hello

import "testing"

// Comment what is *testing.T
// Comment why using *

// *testing.T is a type that allows you to control your test execution.
// using * in a reference means that you are using a pointer to the type.
func TestHello(t *testing.T) {

	want := "Hello World."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
