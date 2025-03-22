package main

import (
	"fmt"
	"strconv"
	"strings"
)



func Exo1(low int , high int) string{
	var tab []string
	
	for i:=low ; i <= high ; i++{
		
		if (i % 7 == 0 && i % 5 != 0){
			fmt.Printf("%d: ", i)
			tab = append(tab, strconv.Itoa(i))
		}

	}

	return strings.Join(tab , ",")

}
func main(){
	Exo1(100 , 200)
}