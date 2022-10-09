package main

import "fmt"

func main()  {
	nums := []int{1,2,2,4,1}
	val := 2
    fmt.Print(removeElement(nums,val))
}

func removeElement(nums []int, val int) int {
    n:= len(nums)
    j:=0

    for i := 0; i < n; i++ {
        if nums[i]!=val{
            nums[j]=nums[i]
            j++
        }
    }
    fmt.Println(nums)
    return j
}