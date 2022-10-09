package main

import "fmt"

func main(){
	nums := []int{2,2,1,1}
	fmt.Print(singleNumber(nums))
}


func singleNumber(nums []int) int {
    temp := make(map[int]int)
	res := 0

    for _,value := range nums{
		temp[value]++
		if temp[value]==1 {
			res += value
		}else{
			res -= value
		}
    }
	return res
}