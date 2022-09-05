package test

import (
	"testing"
)

func Expect(t *testing.T, ok bool) {
	if !ok {
		t.Errorf("oops!")
	}
}
