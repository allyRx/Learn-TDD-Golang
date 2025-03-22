package main

import "testing"

func Testfunc(t *testing.T){
	got := Exo1(100 , 200)
	want := "112,119,126,133,147,154,161,168,182,189,196"

	if got != want {
		t.Errorf("want , %s got %s", want ,got)
	}
}