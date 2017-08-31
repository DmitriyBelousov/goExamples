// - не требует явного указания наследования
// - абстрактны, нельзя инициализировать
// - интерфейс можно встроить в интерфейс
// - в струкутру нельзя встроить интерфейс
//- по соглашению имя интерфейса "жедательно" должно заканчиваться на "er"



package main

import "fmt"

type person struct {
	fname string
	lname string
	age int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	if sa.licenseToKill == true{
		fmt.Println(sa.fname, sa.lname, `says, "I'll kill you"`)
	}else{
		fmt.Println(sa.fname, sa.lname, `says, "Hi"`)
	}
}

func (p person) howOld() {
	fmt.Println(p.fname, p.lname, " - ", p.age, " years old")
}

func (sa secretAgent) howOld() {
	fmt.Println(sa.fname, sa.lname, `says, "I cant't tell you"`)
}

type Speaker interface {
	speak()
}

func saySomething(s Speaker) {
	s.speak()
}


/// инткрфейс встроенный в интерфейс

//type Informater interface {
//	Speaker
//	howOld()
//}

func getInfo(p person, sa secretAgent){
	saySomething(p)
	saySomething(sa)
	p.howOld()
	sa.howOld()
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
		30,
	}

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
			35,
		},
		true,
	}

	saySomething(p1)
	saySomething(sa1)
	getInfo(p1, sa1)
}
