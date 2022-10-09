package main

import "fmt"

func main(){
	accounts := [][]int{{1,2,3},{2,3,4}}
	fmt.Print(maximumWealth(accounts))
}
func maximumWealth(accounts [][]int) int {
	var max int = 0
    for i := 0; i < len(accounts); i++ {
		var sum int = 0
		for j := 0; j < len(accounts[0]); j++ {
			sum = sum + accounts[i][j]
		}
		if sum>max{
			max = sum
		}
	}
	return max
}