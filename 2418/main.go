package main

import (
	"fmt"
	"sort"
)

func main(){
	names:= []string{"neo","ana","anshi"}
	heights:= []int{6,4,5}
	fmt.Print(sortPeople(names,heights))
}
func sortPeople(names []string, heights []int) []string {
	dict:= make(map[int]string)
	ans:= []string{}
    for i,val:= range heights{
		dict[val] = names[i]
	}
	sort.Ints(heights)
	fmt.Println(heights)
	for i:= range heights{
		ans =append(ans, dict[heights[len(heights)-i-1]])
	}
	return ans
}