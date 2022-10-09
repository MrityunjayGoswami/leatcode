package main

import "fmt"

func main(){
	nums :=[]int {2,7,11,15}
	var target int= 9
	fmt.Print(twoSum(nums,target))
}
func twoSum(nums []int, target int) []int {
	result := []int{}
	var l int = len(nums)

    for i:=0; i<l-1; i++ {
		for j:=i+1; j<l; j++{
			if nums[i]+nums[j]==target{
				result = append(result,i,j)
			}
		}
		
	}
	return result
}