package main

import "fmt"

func main()  {
	nums := []int{1,2,4,5,6,7}
	target := 8
    fmt.Print(searchInsert(nums,target))
}
func searchInsert(nums []int, target int) int {
    left := 0
	right := len(nums)
	pivot := 0
	for left<right {
		pivot = (left+right)/2
		if nums[pivot] == target{
			return pivot
		}else if nums[pivot]>target{
			right = pivot
		}else{
			left = pivot+1
		}
	}
	if nums[pivot]>target{
		return pivot
	}else{
		return pivot+1
	}
}