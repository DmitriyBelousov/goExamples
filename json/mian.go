package main

import (
	"encoding/json"
	"fmt"
)

func CheckError(e error){
	if e != nil{
		fmt.Println("ошибка парсинга")
	}
}

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}


func main() {
	//структуры2
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%T, %v\n", res, res)
	fmt.Println(res.Fruits[0])

	//структуры

	//res1D := &Response1{
	//	Page:   1,
	//	Fruits: []string{"apple", "peach", "pear"}}
	//res1B, _ := json.Marshal(res1D)
	////res1B, _ := json.MarshalIndent(res1D, "", "   ")
	//fmt.Println(string(res1B))
	//
	//res2D := &Response2{
	//	Page:   1,
	//	Fruits: []string{"apple", "peach", "pear"}}
	//res2B, _ := json.Marshal(res2D)
	//fmt.Println(string(res2B))

//простые данные
//	bolB, _ := json.Marshal(true)
//	fmt.Println(bolB)
//	fmt.Println(string(bolB))
//
//	intB, _ := json.Marshal(1)
//	fmt.Println(string(intB))
//
//	fltB, _ := json.Marshal(2.34)
//	fmt.Println(string(fltB))
//
//	strB, _ := json.Marshal("gopher")
//	fmt.Println(string(strB))
//
//	slcD := []string{"apple", "peach", "pear"}
//	fmt.Println(slcD)
//	slcB, _ := json.Marshal(slcD)
//	fmt.Println(string(slcB))
//
//	mapD := map[string]int{"apple": 5, "lettuce": 7}
//	mapB, _ := json.Marshal(mapD)
//	fmt.Println(string(mapB))

}



