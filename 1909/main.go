package main

import "fmt"

func main(){
	nums := []int{105,924,32,968}
	fmt.Print(canBeIncreasing(nums))
}
func canBeIncreasing(nums []int)bool{
	var cnt int= 0
    for i := 1; i < len(nums) && cnt < 2; i++ {
        if (nums[i - 1] >= nums[i]) {
            cnt++
            if i > 1 && nums[i - 2] >= nums[i]{
                nums[i] = nums[i - 1]
			}
        }
    }
    return cnt < 2;
}
