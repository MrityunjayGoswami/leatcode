package main

import "fmt"

func main(){
	numRows:= 3
	fmt.Print(generate(numRows))
}
func generate(numRows int) []int {
		arr:=make([][]int,numRows+1)  //make a 2d array
		for i:=0;i<numRows+1;i++{
			arr[i]=make([]int,i+1)      //size of each row will be equal to the row number
		}
		for i:=0;i<numRows+1;i++{
			for j:=0;j<len(arr[i]);j++{
				if (j==0 || j==len(arr[i])-1){
					arr[i][j]=1     //if the element is first or last then it should be 1
				}else{
					arr[i][j]=arr[i-1][j-1]+arr[i-1][j]   //sum of values directly above and to the left of the current element
				}
			}
		}
		return arr[numRows]
	}
