package main

import "fmt"

func main(){
	var num []int
	var arrsize int
	var temp int
	fmt.Print("Enter the size of the array : ")
	fmt.Scanln(&arrsize)

	for i:=0; i<arrsize; i++{
		fmt.Printf("Enter the %dth element of array : ",i+1)
		fmt.Scanln(&temp)
		num = append(num, temp)
	}
	fmt.Println("Entered array is : ",num)

	var ans []int
	ans = getConcatenation(num)
	fmt.Println(ans)
}


func getConcatenation(slice []int) []int {
    var n int = len(slice)
	var ans = []int{}

	for i := 0; i < n; i++ {
		ans = append(ans,slice[i])
	}

	for i := n; i < 2*n; i++ {
		ans =append(ans,slice[i-n])
	}

	return ans
	
}