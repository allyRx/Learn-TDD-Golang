package main

import (
	"bytes"
	"fmt"
)



func Greet(writer *bytes.Buffer , name string){
	fmt.Printf("Hello, %s", name)
}
func main(){
	fmt.Println("hello")
}