package main

import "fmt"

func main(){
	strs := []string{"flower","flower","flwer"}
	fmt.Print(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs)==0{return ""}
	if len(strs)==1{return strs[0]}
	mx := 0

	for{
		for i:=1; i<len(strs); i++{
			if mx >= len(strs[i-1]) || mx >= len(strs[i]) || strs[i-1][mx]!= strs[i][mx] {
				return strs[0][:mx]
			}
		}
		mx++
	}
	
}