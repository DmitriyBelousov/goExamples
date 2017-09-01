

// выбор по значению

// switch необязательнаяИнструкция; необязательноеВыражение {
// case списокВыражений1: блок1
// ...
// case списокВыраженийN: блокN
// default

// Точка с запятой обязательна при наличии необязательной инструкции, независимо от наличия необязательного выражения. 
// Если в инструкции switch отсутствует необязательное выражение,компилятор предполагает, что в качестве выражения используется
// значение true. 
//  Поскольку выполнение автоматически не«проваливается» через границы разделов case, нет необходимости
// вставлять инструкцию break в конец каждого блока case. Если такое «проваливание» желательно, достаточно просто добавить инструк-
// цию fallthrough. 

package main

import (
	"time"
	"fmt"
	"strings"
	"path/filepath"
)

func main(){
	switchEx3()

	switchEx("testfile.gz")
	switchEx2("blabla.zip")

	//i := 2
	//fmt.Print("Write ", i, " as ")
	//switch i {
	//case 1:
	//	fmt.Println("one")
	//case 2:
	//	fmt.Println("two")
	//case 3:
	//	fmt.Println("three")
	//}

}

func switchEx3(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}


func switchEx(file string){
	switch suffix := Suffix(file); suffix { 
	case ".gz":
		GzipFileList()
	case ".tar":
		fallthrough
	case ".tar.gz":
		fallthrough
	case ".tgz":
		TarFileList()
	case ".zip":
		ZipFileList()
	}
}

func switchEx2(file string){
	switch Suffix(file) {
	case ".gz":
		GzipFileList()
	case ".tar", ".tar.gz", ".tgz":
		TarFileList()
	case ".zip":
		ZipFileList()
	}
}
func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
	return file[i:]
	}
	return file
}

func GzipFileList(){
	fmt.Println(".gz")
}

func TarFileList(){
	fmt.Println(".tar")
}

func ZipFileList(){
	fmt.Println(".zip")
}

