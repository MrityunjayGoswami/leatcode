package main

import "fmt"

func main(){
	nums := []int{3,2,3}
	fmt.Print(majorityElement(nums))
}

func majorityElement(nums []int) int {
    n := len(nums)/2
	count := make(map[int]int)

	for _,val := range nums{
		count[val]++
		if count[val]>n {
			return val
		}
	}
	return -1
}