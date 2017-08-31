package main

import "fmt"

type Man struct{
	Name string
	Age int
}


func main(){

	x := 15
	fmt.Println(x, &x )

	//y := &x
	//fmt.Println(y)
	//
	//z := &y
	//fmt.Println(*z)

	slice := []int{1,2,3,4,5}
	fmt.Println(slice)
	change2(slice)
	fmt.Println(slice)

	modifySlice := change(slice)
	fmt.Println(modifySlice)
	fmt.Println("-------------------")

	man1 := Man{ "Bob", 23}
	fmt.Println(man1)

	changeStruct(man1)
	fmt.Println(man1)

	changeStruct2(&man1)
	fmt.Println(man1)

	//i := 1
	//fmt.Println(i)
	//setZero(i)
	//fmt.Println(i)
	//
	//zeroPointer(&i)
	//fmt.Println(i)

}

func setZero(val int){
	val = 0
}

func zeroPointer(val *int){
	*val = 0
}

func change(ar []int)[]int{
	ar[0] = 99
	return ar
}

func change2(ar []int){
	ar[0] = 99
}

func changeStruct(man Man){
	man.Name = "blabla"
}

func changeStruct2(man *Man){
	man.Name = "blabla"
}
