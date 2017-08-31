package main

import (
	"fmt"
	"reflect"
)

var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}
var slice = make([]int, 0)
var intarray = []int{1, 2, 4, 8, 16}
var mapone = map[int]string{}
var maptwo = map[string]interface{}{}

func main() {

	for i := 0; i < 5; i++ {

		// print the $th value of the intarray and the strarray
		//fmt.Println(intarray[i], "\t", strarray[i])

		mapone[intarray[i]] = strarray[i]
		maptwo[strarray[i]] = mapone
	}
	//fmt.Println(mapone)
	//fmt.Println(maptwo)

	tst := "string"
	tst2 := 10
	tst3 := 1.2

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))
	fmt.Println(reflect.TypeOf(strarray))
	fmt.Println(reflect.TypeOf(slice))

}