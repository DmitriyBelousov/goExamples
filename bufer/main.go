package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var buffer bytes.Buffer
	for i := 0; i < 2; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer)
	fmt.Printf("%T, %v\n", buffer, buffer)


	s := buffer.String()
	fmt.Println(s)

//конвертирование числа в строку
	for i := 0; i < 8; i++ {
		s += strconv.Itoa(i)
	}
	fmt.Printf("%T, %v\n", s, s)



	s += buffer.String()
	fmt.Printf("%T, %v\n", s, s)

	for i := 0; i < 5; i++ {
		s = "pre" + s
	}
	fmt.Printf("%T, %v\n", s, s)


}