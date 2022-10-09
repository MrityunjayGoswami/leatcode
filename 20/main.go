package main

import "fmt"

func main(){
	s := "}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
    stack := make([]rune, 0)

	for _, i:= range s{
		if i == '(' || i =='{' || i == '['{
			stack = append(stack,i)
		}else{
			top := len(stack)
			if len(stack)==0 || stack[top-1] == '(' && i!= ')' || stack[top-1]=='{' && i!='}' || stack[top-1]=='[' && i!=']' {
				return false
			}
			if len(stack)>0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack)==0
}