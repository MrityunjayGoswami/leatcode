package main

import "fmt"

func main(){
	nums := []int{1,2,2,1}
	var k int = 1
	fmt.Print(countKDifference(nums,k))
}

func countKDifference(nums []int, k int) int {
	var result int =  0

	for i := range nums{
		for j := range nums{
			if nums[i]-nums[j] == k {
				result = result+1

			}
		}
	}
	return result
}