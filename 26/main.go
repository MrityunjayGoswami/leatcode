package main

import "fmt"

func main(){
	nums :=  []int{1,1,1,2,2,3}
	fmt.Print(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	left := 0
	ln := len(nums)
	if ln <= 1 {
        return ln
    }

	for right:=1; right<ln; right++{
		if nums[left]!=nums[right]{
			left++
			nums[left] = nums[right]
		}
	}
	return left+1
}