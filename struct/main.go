package main

import "math"


type Rectangle struct{
	Width float64
	Height float64
}

type Cercle struct {
	Rayon float64
}

func Perimetre(size Rectangle) float64{
	return (size.Height + size.Width) * 2
}

func (size Rectangle) Aires() float64{
	return size.Height * size.Width
}

func (size Cercle) Aires() float64{
	return math.Pi * size.Rayon * size.Rayon
}