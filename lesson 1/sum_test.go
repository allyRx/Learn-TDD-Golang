package main

import "testing"

func TestSum(t *testing.T){
	Sum := Sum(4,3)
	expected := 7

	if Sum != expected {
		t.Errorf("got %q want %q", Sum , expected)
	}
}