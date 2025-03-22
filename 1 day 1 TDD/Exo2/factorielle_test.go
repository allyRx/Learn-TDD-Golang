package main

import "testing"



func TestFactorielle(t *testing.T){
	
	IsEquals := func(t testing.TB , got int, want int) {
		if got != want {
			 t.Errorf("expected: %d , want : %d" , got , want)
		}
	}
	
	got := fact(8)
	want := 40320

	IsEquals(t ,got , want)
	
	
	got1 := fact(0)
	want1 := 1

	IsEquals(t, got1 , want1)

}