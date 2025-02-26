package main



func SumArray(numbers []int) int{
	var result int
	
	for _, number := range numbers {
		result += number
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

func SumAllTails(numbersToSum ...[]int) []int{
	
	var sum []int

	for _, result:=range numbersToSum{
		if len(result) == 0 {
			sum = append(sum, 0)
			continue
		}else{
			tail := result[1:]
			sum = append(sum, SumArray(tail))
		}
		
	}

	return sum
}

func main(){
}

