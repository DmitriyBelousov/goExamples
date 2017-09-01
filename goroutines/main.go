package main

import (
	"fmt"
	"sync"
)


func main(){
	wg := new(sync.WaitGroup)

	 // go showString("bla")
	 //go showString("gogo")

	for i:=0; i < 10; i++{
		go showString2("ttt", wg)
		go showString2("0000", wg)
	}

	wg.Wait()
	//fmt.Scanln()
}

func showString(someStr string){
	for i := 0; i < 10; i++{
		fmt.Println( i, " - ", someStr)
	}
}

func showString2(someStr string, wg *sync.WaitGroup){
	wg.Add(1)
	defer wg.Done()
	fmt.Println(someStr)
}