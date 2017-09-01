package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Producer struct{
	OutChan chan int
}

func (p *Producer) getOutChan() <-chan int{
	return p.OutChan
}

func (p *Producer) produce(){
	for{
		time.Sleep(1 * time.Second)
		p.OutChan <- rand.Int()
	}
}

func main(){

	prod := Producer{make(chan int)}
	go prod.produce()
	//for i:= range prod.getOutChan(){
	//	fmt.Println("Message from chan", i)
	//	prod.getOutChan() <- 3
	//
	//}
	for{
		i:= <- prod.getOutChan()
		fmt.Println("Message from chan", i)
		prod.getOutChan() <- 3

	}

}
