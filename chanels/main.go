package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Producer struct{
	OutChan chan int
}

func (p *Producer) produce(){
	for{
		time.Sleep(1 * time.Second)
		p.OutChan <- rand.Int()
	}
}

func main(){

	//intChan := make(chan int,10)
	//intChan <- 1
	//fmt.Println(<-intChan)
	readChan := make(chan int)
	prod := Producer{readChan}
	go prod.produce()
	for{
		i:= <-readChan
		fmt.Println("Message from chan", i)
	}

}
