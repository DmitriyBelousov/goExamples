package main

import "fmt"

func main(){

	a,b := vals()
	fmt.Println(a,b)

	_,c := vals()
	fmt.Println(c)
	fmt.Println("----------------------")

	sum(1,2,3)
	sum(1,2,5,56,23,900)
}

func vals()(int,int){
	return 3,7
}

func sum(nums ... int){
	fmt.Println(nums)
	total := 0
	for _, num := range  nums{
		total += num
	}
	fmt.Println(total)
}
