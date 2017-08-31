package main

import (
	"fmt"
)

type Day int

const (
	SUNDAY Day = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

const(
	a = iota *2
	b
	c
	d
	f
)


func main() {
	fmt.Printf("Monday has the value %d\n", MONDAY)
	fmt.Printf("Friday has the value %d\n", FRIDAY)
	fmt.Println(a,b,c,d,f)
}