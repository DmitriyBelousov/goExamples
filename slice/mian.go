package main


import(
	"fmt"
)

func main(){

	first := make([]byte, 5,10)
	second := make([]byte,3)
	var third []bool
	fourth := []int{1,2,3,4,5}

	fmt.Printf("first len(%d) cap(%d) = %v\n", len(first),cap(first), first)
	fmt.Printf("second len(%d) cap(%d) = %v\n", len(second),cap(second), second)
	fmt.Printf("third len(%d) cap(%d) = %v\n", len(third),cap(third), third)
	fmt.Printf("fourth len(%d) cap(%d) = %v\n", len(fourth),cap(fourth), fourth)

	fmt.Println("----------------------------------------")

	part1 := fourth[:2]
	part2 := fourth[2:4]
	part3 := fourth[4:]
	part4 := fourth[:]

	fmt.Printf("part1 len(%d) cap(%d) = %v\n", len(part1),cap(part1), part1)
	fmt.Printf("part2 len(%d) cap(%d) = %v\n", len(part2),cap(part2), part2)
	fmt.Printf("part3 len(%d) cap(%d) = %v\n", len(part3),cap(part3), part3)
	fmt.Printf("part4 len(%d) cap(%d) = %v\n", len(part4),cap(part4), part4)

	part1[0] = 100
	part4[4] = 90
	fmt.Println(fourth)
}