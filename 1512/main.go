package main

import (
	"fmt"
)
func main()  {
	nums := []int{1,2,3,1,1,3}
	fmt.Print(numIdenticalPairs(nums))
}
func numIdenticalPairs(nums []int) int {
    c:=0
    m := make([]int,101)
    for i:=0;i<len(nums);i++{
            c+=m[nums[i]]
			fmt.Println(c,m[nums[i]])
            m[nums[i]]++
    }
    return c
}