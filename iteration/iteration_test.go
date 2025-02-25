package main

import "testing"

func TestIteration(t *testing.T){
	Iteration := Iteration("y" , 2) 
	expected := "yy"

	if Iteration != expected{
		t.Errorf("got %q want %q", Iteration, expected)
	}
}