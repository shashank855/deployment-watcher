package main

import (
	"testing"

	"github.com/shashank855/recursion/fibanacci/fibanacci"
)

func TestFib(t *testing.T) {
	got := fibanacci.Fib(6)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
