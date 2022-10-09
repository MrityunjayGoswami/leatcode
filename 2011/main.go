package main

import "fmt"

func main(){

	operations := []string{"X++","++X","X--"}
	fmt.Print(finalValueAfterOperations(operations))

}

func finalValueAfterOperations(operations []string) int {

	var result int = 0 
	for i := range operations{
		if operations[i] == "X++"|| operations[i] == "++X" {
			result = result+1
		
		}else if operations[i] == "--X" || operations[i] == "X--"{
			result = result-1
		}
	}
	return result
}