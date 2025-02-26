package main

import "math"


type Rectangle struct{
	Width float64
	Height float64
}

type Cercle struct {
	Rayon float64
}

type Shape interface{
	Aires() float64
	Perimetre() float64
}



func (size Rectangle) Perimetre() float64{
	return (size.Height + size.Width) * 2
}

func (size Cercle) Perimetre() float64{
	return 2 * math.Pi * size.Rayon
}

func (size Rectangle) Aires() float64{
	return size.Height * size.Width
}

func (size Cercle) Aires() float64{
	return math.Pi * size.Rayon * size.Rayon
}