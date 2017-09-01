package main

import "fmt"

func main(){
	//fmt.Println("hello")
	//panic("df")

	fmt.Println("main.begin")

	defer func(){
		fmt.Println("main.defer")

		if err := recover(); err != nil{
			fmt.Println("main.recover: ", err)
		}
	}()

	first()
	fmt.Println("main.end")

}

func first(){
	fmt.Println("first.begin")

	defer func(){
		fmt.Println("first.defer")
		if err := recover(); err !=nil{
			fmt.Println("first.recover: ", err)
			//panic(err)
		}
	}()

	second()

	fmt.Println("first.end")
}

func second(){
	fmt.Println("second.begin")

	defer func(){
		fmt.Println("second.defer")

		if err := recover(); err != nil{
			fmt.Println("second.recover: ", err)
			panic(err)
		}
	}()

	panic("Что-то пошло не так")
	fmt.Println("second.end")
}