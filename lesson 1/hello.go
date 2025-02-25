package main

import "fmt"


func Hello() string {
	return "Hello , Word!"
}


func Sum(a int, b int) int{
	return a + b
}

func Soustraction(a int , b int) int{
	return a - b
}

func main(){
	fmt.Println(Hello())
	fmt.Println(Sum(4,3))
}