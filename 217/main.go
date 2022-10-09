package main

import "fmt"

func main(){
	nums := []int{3,2,5}
	fmt.Print(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	count := make(map[int]int)

	for _,val := range nums{
		count[val]++
	}
	for val:= range count{
		if count[val]>1{
			return true
		}
	}
	return false
}