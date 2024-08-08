package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Armaan")

	got := buffer.String()
	want := "Hello Armaan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
