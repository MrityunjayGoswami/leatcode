package main

import "fmt"

func main(){
	//boolean datatype
	b := 1>2

	//array data type
	array1 := [4]int{1,2,3,4}

	//slice data type
	slice1 := []int{}

	//int data type
	int1 := 1

	//string data type
	string1 := "hello"

	//map data type
	map1 := map[int]string{
		1 : "one",
		2 : "two",
	}

	//struct data type
	type struct1 struct{
		name string
		age int
	}
	//pointer data type
	var pointer *int
	pointer = &int1

	//channel data type
	c := make(chan string)

	//interface
	type bot interface{
	}

	hello(array1)
	structure()
	fmt.Println(map1)
}
//funtion data type
func hello(array1){
	fmt.Println(array1)
}
//method data type
func (d struct1) structure(){
	fmt.Println(struct1())
}