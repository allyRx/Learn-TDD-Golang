package main

import (
	"reflect"
	"testing"
)



func TestDouble(t *testing.T){
	input := 8
	want := map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64}
	got := makeDouble(input)

	if !reflect.DeepEqual(got , want){
		t.Errorf("Expected : %d,  want : %d \n", got, want)
	}
}