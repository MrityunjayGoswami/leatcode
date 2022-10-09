package main

import "fmt"

func main(){
	nums := []int{2,5,1,3,4,7}
	var n int = 3
	fmt.Print(shuffle(nums,n))
}
func shuffle(nums []int, n int) []int {
    result :=make([]int,len(nums))
	c := 0

	for i := 0; i < n; i++ {
		result[c] = nums[i]
		c++
		result[c] = nums[n+i]
		c++
	}
	return result
}