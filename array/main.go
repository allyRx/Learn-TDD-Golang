package main

import "fmt"

func SumArray(numbers []int) int{
	var result int
	for _,nb:=range numbers{
		result += nb
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int{
	
	lengthOfNumbers := len(numbersToSum) //verifie le nombre d'array passer par le fontion
	sums := make([]int, lengthOfNumbers) //creer une array vide base sur le nombre d'array passer par la fonction
	
	for i, numbers := range numbersToSum {
		
		sums[i] = SumArray(numbers)
	}

	return sums
}
func main(){
	test := SumAll([]int{1,2}, []int{0,9}, []int{0,9})

	fmt.Println(test)
}