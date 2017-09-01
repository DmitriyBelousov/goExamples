//for { // Бесконечный цикл
//блок
//}
//for логическоеВыражение { // Цикл while
//блок
//}
//for необязательнаяПредИнструкция;логическоеВыражение;необязательнаяПостИнструкция {
//блок
//}
//for индекс, символ := range Строка { // Итерации по символам в Строке
//блок
//}
//for индекс := range Строка { // Итерации по символам в Строке 
//блок // символ, размер := utf8.DecodeRuneInString(Строка[индекс:])
//}
//for индекс, элемент := range массивИлиСрез { // Итерации по массиву или срезу 
//блок
//}
//for index := range массивИлиСрез { // Итерации по массиву или срезу 
//блок // item := массивИлиСрез[index]
//}
//for ключ, значение := range Отображение { // Итерации по элементам отображения 
//блок
//}
//for ключ := range Отображение { // Итерации по элементам отображения 
//блок // значение := Отображение[ключ]
//}
//for элемент := range Канал { // Итерации по элементам в канале
//блок
//}




package main

import "fmt"

func main(){
	//i := 1
	//for i <= 3 {
	//	fmt.Println(i)
	//	i = i + 1
	//}
	//
	//for j := 7; j <= 9; j++ {
	//	fmt.Println(j)
	//}
	//
	//for {
	//	fmt.Println("loop")
	//	break
	//}
	//
	//for n := 0; n <= 5; n++ {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}

	grades := []int{87, 55, 43, 71, 60, 43, 32, 19, 63}
	fmt.Println(grades)
	inflate2(grades, 3)
	fmt.Println(grades)

	inflate(grades, 3)
	fmt.Println(grades)

}

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

func inflate2(numbers []int, factor int) {
	for _, item := range numbers {
		item *= factor
	}
}


