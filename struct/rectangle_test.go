package main

import "testing"

func TestPerimetre(t *testing.T){
	
	// checkPerimetre := func(t testing.TB , shape Shape , want float64){
	// 	got := shape.Perimetre()

	// 	if got != want{
	// 		t.Errorf("got %f want %f", got, want)
	// 	}
	// }
	
	
	// t.Run("Rectangle" , func(t *testing.T){
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	checkPerimetre(t, rectangle, 40.0)
	// })

	// t.Run("Cercle" , func(t *testing.T){
	// 	cercle := Cercle{2}
	// 	checkPerimetre(t, cercle, 12.566370614359172)
	// })

	perimetreTest := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{10.0 , 10.0} , 40.0},
		{Cercle{2} , 12.566370614359172},
	}

	for _,tt := range perimetreTest{
		got := tt.shape.Perimetre()

		if got != tt.want{
			t.Errorf("got %f want %f", got, tt.want)
		}
	}

}


func TestAire(t *testing.T){
	
	aireTest := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Cercle{10}, 314.1592653589793},
		
	}
	
	for _ , tt := range aireTest{
		got := tt.shape.Aires()
		if got != tt.want{
			t.Errorf("got %f want %f", got, tt.want)
		}
	}
	
}