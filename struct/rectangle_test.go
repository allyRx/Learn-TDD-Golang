package main

import "testing"

func TestPerimetre(t *testing.T){
	rectangle := Rectangle{10.0,10.0}
	got := Perimetre(rectangle)
	want := 40.0

	if got != want{
		t.Errorf("got %f want %f", got, want)
	}
}


func TestAire(t *testing.T){
	t.Run("Aire pour le rectangle" , func(t *testing.T){
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Aires()
		want := 72.0

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})


	t.Run("Aire pour le cercle" , func(t *testing.T){
		cercle := Cercle{10}

		got := cercle.Aires()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}