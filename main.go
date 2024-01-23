package main

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

func main(){
	result:=Sum(3,4)
	fmt.Println("Sum is:", result)
}