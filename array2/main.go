package main

import(
	"fmt"
)

func main(){

	var first [2]bool
	second := [...]int {1,-2,3}
	third := [6]uint{1,5,9,6}

	var fourth [2][2]float32
	fifth:= [2][2]int{ {1}, {5,4} }

	fmt.Printf("first len(%d) = %v\n", len(first), first)
	fmt.Printf("second len(%d) = %v\n", len(second), second)
	fmt.Printf("third len(%d) = %v\n", len(third), third)
	fmt.Printf("fourth len(%d) = %v\n", len(fourth), fourth)
	fmt.Printf("fifth len(%d) = %v\n", len(fifth), fifth)

	fmt.Println(second[2])

	copyOfThird := third
	//copyOfThird := &third
	copyOfThird[1] = 23
	copyOfThird[2] = 55
	fmt.Println(third)
	fmt.Println(copyOfThird)

}
