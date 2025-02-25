package main

import "testing"


func TestSoustraction(t *testing.T){
	got := Soustraction(4,3)
	want := 1

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}