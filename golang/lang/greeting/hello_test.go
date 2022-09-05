package greeting

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("Golang")
	expect := "Hello, Golang!"
	if actual != expect {
		t.Errorf("expect: %s, but got %s", expect, actual)
	}
}
