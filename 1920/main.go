package main

import "fmt"

func main(){
	var nums = []int {0,2,1,5,3,4}
	fmt.Print(buildArray(nums))
}
func buildArray(nums []int) []int {
	var result = []int{}

    for i := 0; i < len(nums); i++ {
		result = append(result,nums[nums[i]]) 
	}
	return result
}