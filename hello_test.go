package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Taufik")
	want := "Hello, Taufik"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}