package main

import (
	"fmt"
	"sort"
)

func main(){
	nums1 := []int{1,2,3}
	m:=3
	nums2 := []int{2,7,8,9}
	n:=4
	merge(nums1,m,nums2,n)
}
func merge(nums1 []int, m int, nums2 []int, n int)  {
	ans := []int{}
    ans= append(nums1[:m],nums2[:n]...)
	sort.Ints(ans)
	fmt.Println(ans)
}