package main

import "testing"



func TestSummArray(t *testing.T){
	
	t.Run("Collection of 5 Numbers" , func(t *testing.T){
		numbers := []int {1,2,3,4,5}
	
	got := SumArray(numbers)
	want := 15


	if got != want {
		t.Errorf("got %d want %d", got , want)
	}
	})

	t.Run("Collection any numbers" , func(t *testing.T) {
		number := []int{1,3}

		got := SumArray(number)
		want := 4

		if got != want {
			t.Errorf("got %d want %d" , got , want)
		}

	})
	

}

//SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}
func TestSumAll(t *testing.T){
	
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3, 9}

	if got != want {
		t.Errorf("Got %v , want %v" , got , want)
	}
}