package main

import "fmt"

func main(){
	x := 1214
	fmt.Print(isPalindrome(x))

}

func isPalindrome(x int) bool {
	res := []int{}
	if x<0{
		return false
	}

	for x!=0{
		res = append(res, x%10)
		x /= 10
	}
	for i:=0; i<len(res)/2; i++{
		if res[i]!=res[len(res)-1-i]{
			return false
		}
	}
	return true
}