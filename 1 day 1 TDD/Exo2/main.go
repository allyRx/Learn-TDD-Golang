package main

import (
	"fmt"
	"log"
)

func giveNumber() int{
	var input int

	fmt.Print("Entrez le nombre que vous cherchez son factorielle: ")
	_ , err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("S'il vous plait entrez un nombre: ")
	}

	return input
}

func fact(n int) int{
	result := 1
	
	if n < 0 {
		
		return 0
	}
	
	for i := 1; i <= n ; i++ {
		result = result * i
	}
	return result
}

func main(){
	input := giveNumber()
	fmt.Println(input)
	fmt.Printf("La factorielle de  %d est %d \n",input , fact(input))
}