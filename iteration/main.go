package main

import "fmt"

func Iteration(a string , nb int16) string{
	
	var repeated string

	for range nb{
		repeated += a
	}

	return repeated
}


func main(){
	result := Iteration("a" , 8)
	fmt.Println(result)
}