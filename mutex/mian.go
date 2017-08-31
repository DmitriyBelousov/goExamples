package main

import (
	"fmt"
	"sync"
)

type MyObj struct{
	MapMutex 	*sync.Mutex
	Map 		map[string]string
}

func main(){

	obj := MyObj{
		Map:make(map[string]string,0),
		MapMutex: new(sync.Mutex),
	}
	for i:=0; i <1000; i++{
		go writeToMap(&obj)

	}


	fmt.Scanln()
}

func writeToMap(obj *MyObj){
	obj.MapMutex.Lock()
	defer obj.MapMutex.Unlock()
	obj.Map["hello"] = "123"
	fmt.Println(obj.Map)

}
