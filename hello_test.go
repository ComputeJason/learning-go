package main

import (
	"testing"
)

func Test_Hello(t *testing.T){

	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}