package main

import "fmt"



func makeDouble(n int) map[int] int{
	var result = make(map[int] int , n)
	
	for i := 0 ; i <= n ; i++{
		result[i] = i * i
	}
	
	return result
}
func main(){
	fmt.Println(makeDouble(8))
}