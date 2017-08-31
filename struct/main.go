package main

import "fmt"

type Man struct{
	Name string
	Age uint
}

type Employee struct{
	Person Man
	Salary uint
}

func main(){

	var man1 = Man{}
	fmt.Println(man1)

	var man2 = new(Man)
	fmt.Println(man2)
	fmt.Println(&man2)

	man3 := Man{"Bob", 90}
	fmt.Println(man3.Name)

	man3.bDay()
	fmt.Println(man3)

	//man4 := &man3
	//man4.Age = 100
	//fmt.Println("man4 - ", *man4)
	//fmt.Println("man3 - " ,man3)

	//
	//emp := new(Employee)
	//emp.Person.bDay()
	//fmt.Println(*emp)

}

func(man *Man) bDay(){
	man.Age++
}