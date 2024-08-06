package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Armaan");
	want:= "Hello Armaan."
	
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}