package main

import "fmt"

func main(){
	prices := []int{7,1,5,3,6,4}
	fmt.Print(maxProfit(prices))
}

func maxProfit(prices []int) int {
    max,lowest := 0,0
	for i := 1; i < len(prices); i++ {
		if prices[i]<prices[lowest]{
			lowest = i
		}else if cur:=prices[i]-prices[lowest]; cur>max{
			max = cur
		}
	}
	return max
}