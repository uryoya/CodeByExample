package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := sum(1, 1)
	expected := 2
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
